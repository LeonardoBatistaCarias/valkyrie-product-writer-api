package product_processor

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/segmentio/kafka-go"
	"sync"
)

type ProductMessageProcessor struct {
	cfg      *config.Config
	commands commands.ProductCommands
}

func NewProductMessageProcessor(cfg *config.Config, commands commands.ProductCommands) *ProductMessageProcessor {
	return &ProductMessageProcessor{cfg: cfg, commands: commands}
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
			fmt.Errorf("workerID: %v, err: %v", workerID, err)
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
	//s.metrics.SuccessKafkaMessages.Inc()
	//s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		//s.log.WarnMsg("commitMessage", err)
	}
}

func (p *ProductMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	//s.metrics.ErrorKafkaMessages.Inc()
	//s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		//s.log.WarnMsg("commitMessage", err)
	}
}

func (p *ProductMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	//s.log.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}
