package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize = 30
)

type productMessageProcessor struct {
	cfg      *config.Config
	commands commands.ProductCommands
}

func NewProductMessageProcessor(cfg *config.Config, commands commands.ProductCommands) *productMessageProcessor {
	return &productMessageProcessor{cfg: cfg, commands: commands}
}

func (s *productMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
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
		case s.cfg.KafkaTopics.ProductCreate.TopicName:
			s.processCreateProduct(ctx, r, m)
		case s.cfg.KafkaTopics.ProductDelete.TopicName:
			s.processDeleteProductByID(ctx, r, m)
		case s.cfg.KafkaTopics.ProductUpdate.TopicName:
			s.processUpdateProductByID(ctx, r, m)
		}
	}
}
