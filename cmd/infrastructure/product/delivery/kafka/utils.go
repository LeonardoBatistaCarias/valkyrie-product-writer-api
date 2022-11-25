package kafka

import (
	"context"

	"github.com/pkg/errors"

	"github.com/segmentio/kafka-go"
)

func (s *productMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	if err := r.CommitMessages(ctx, m); err != nil {
		errors.Errorf("commitMessage", err)
	}
}

func (s *productMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	if err := r.CommitMessages(ctx, m); err != nil {
		errors.Errorf("commitMessage", err)
	}
}
