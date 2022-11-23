package kafka

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/create"
	kafkaMessages "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/proto/kafka"
	"github.com/avast/retry-go"
	uuid "github.com/satori/go.uuid"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *productMessageProcessor) processCreateProduct(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &kafkaMessages.ProductCreate{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		fmt.Errorf("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	command := create.NewCreateProductCommand(msg.GetName(), msg.GetDescription(), 1, msg.GetPrice(), msg.GetQuantity(), uuid.NewV4(), nil, true)

	if err := retry.Do(func() error {
		return s.commands.CreateProduct.Handle(ctx, *command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		fmt.Errorf("CreateProduct.Handle", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
