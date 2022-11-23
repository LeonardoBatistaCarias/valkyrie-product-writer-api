package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	kafkaClient "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/postgres"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product"
	kafkaConsumer "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product/delivery/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	cfg       *config.Config
	kafkaConn *kafka.Conn
	pgConn    *pgxpool.Pool
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	pgxConn, err := postgres.NewPgxConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.pgConn = pgxConn
	log.Printf("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()

	productRepo := repository.NewProductRepository(s.cfg, pgxConn)
	pgGateway := product.NewProductPostgresGateway(productRepo)
	createProductCommand := create.NewCreateProductHandler(pgGateway)
	productCommands := commands.NewProductCommands(createProductCommand)
	productMessageProcessor := kafkaConsumer.NewProductMessageProcessor(s.cfg, *productCommands)

	log.Println("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), kafkaConsumer.PoolSize, productMessageProcessor.ProcessMessages)

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errcheck

	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}

	s.runHealthCheck(ctx)
	s.runMetrics(cancel)

	<-ctx.Done()

	return nil
}
