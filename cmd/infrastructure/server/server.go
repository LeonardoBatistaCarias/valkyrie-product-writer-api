package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	kafkaConsumer "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor/product_processor"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/metrics"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product/persistence"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/logger"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/postgres"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type server struct {
	log       logger.Logger
	cfg       *config.Config
	kafkaConn *kafka.Conn
	pgConn    *pgxpool.Pool
	m         *metrics.Metrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.m = metrics.NewMetrics(s.cfg)

	pgxConn, err := postgres.NewPgxConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.pgConn = pgxConn
	s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	productRepo := persistence.NewProductRepository(s.cfg, pgxConn)
	pgGateway := product.NewProductPostgresGateway(productRepo)
	productCommands := commands.NewProductCommands(pgGateway)
	rc, err := grpc.NewReaderServiceClient(ctx, s.cfg)
	if err != nil {
		s.log.Errorf("Error in connecting grpc reader service PORT ", err)
		return err
	}
	productMessageProcessor := kafkaConsumer.NewProductMessageProcessor(s.cfg, *productCommands, rc, s.m, s.log)

	if err := s.initKafka(ctx, productMessageProcessor); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (s *server) initKafka(ctx context.Context, processor message_processor.MessageProcessor) error {
	log.Println("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumer(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID)
	go cg.ConsumeTopics(ctx, s.getConsumerGroupTopics(), constants.POOL_SIZE, processor.ProcessMessage)

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close()

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}
	return nil
}
