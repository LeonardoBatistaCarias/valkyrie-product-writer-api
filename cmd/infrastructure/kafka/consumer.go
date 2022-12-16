package kafka

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type Consumer struct {
	Brokers []string
	GroupID string
}

func NewConsumer(brokers []string, groupID string) *Consumer {
	return &Consumer{Brokers: brokers, GroupID: groupID}
}

func (c *Consumer) ConsumeTopics(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	r := newKafkaReaderTest(c.Brokers, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("consumerGroup.r.Close: %v", err)
		}
	}()

	log.Printf("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}

func newKafkaReaderTest(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               constants.MIN_BYTES,
		MaxBytes:               constants.MAX_BYTES,
		QueueCapacity:          constants.QUEUE_CAPACITY,
		HeartbeatInterval:      constants.HEART_BEAT_INTERVAL,
		CommitInterval:         constants.COMMIT_INTERVAL,
		PartitionWatchInterval: constants.PARTITION_WATCH_INTERVAL,
		MaxAttempts:            constants.MAX_ATTEMPTS,
		MaxWait:                constants.MAX_WAIT,
		Dialer: &kafka.Dialer{
			Timeout: constants.DIAL_TIMEOUT,
		},
	})
}
