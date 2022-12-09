package kafka

import (
	"context"
	"github.com/avast/retry-go"
	"time"

	"github.com/pkg/errors"

	"github.com/segmentio/kafka-go"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
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
