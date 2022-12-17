package product_processor

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/pb/reader_service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/metrics"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/logger"
	"github.com/segmentio/kafka-go"
	"sync"
)

type ProductMessageProcessor struct {
	cfg      *config.Config
	commands commands.ProductCommands
	rc       reader_service.ProductReaderServiceClient
	metrics  *metrics.Metrics
	log      logger.Logger
}

func NewProductMessageProcessor(cfg *config.Config, commands commands.ProductCommands, rc reader_service.ProductReaderServiceClient, metrics *metrics.Metrics, log logger.Logger) *ProductMessageProcessor {
	return &ProductMessageProcessor{cfg: cfg, commands: commands, rc: rc, metrics: metrics, log: log}
}

func (p *ProductMessageProcessor) ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		m, err := r.FetchMessage(ctx)

		if err != nil {
			p.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		switch m.Topic {
		case p.cfg.KafkaTopics.ProductCreate.TopicName:
			p.processCreateProduct(ctx, r, m)
		case p.cfg.KafkaTopics.ProductDelete.TopicName:
			p.processDeleteProductByID(ctx, r, m)
		case p.cfg.KafkaTopics.ProductDeactivate.TopicName:
			p.processDeactivateProductByID(ctx, r, m)
		case p.cfg.KafkaTopics.ProductUpdate.TopicName:
			p.processUpdateProductByID(ctx, r, m)
		}
	}
}

func (p *ProductMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	p.metrics.SuccessKafkaMessages.Inc()
	p.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		p.log.WarnMsg("commitMessage", err)
	}
}

func (p *ProductMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	p.metrics.ErrorKafkaMessages.Inc()
	p.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		p.log.WarnMsg("commitMessage", err)
	}
}

func (p *ProductMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	p.log.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}
