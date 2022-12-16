package message_processor

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"sync"
)

var (
	RETRY_OPTIONS = []retry.Option{retry.Attempts(constants.RETRY_ATTEMPTS), retry.Delay(constants.RETRY_DELAY), retry.DelayType(retry.BackOffDelay)}
)

type MessageProcessor interface {
	ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}
