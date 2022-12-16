package reader_service

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	reader_service "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/reader_service/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/constants"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func NewReaderServiceClient(ctx context.Context, cfg *config.Config) (reader_service.ProductReaderServiceClient, error) {
	readerServiceConn, err := newReaderServiceConn(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return reader_service.NewProductReaderServiceClient(readerServiceConn), nil
}

func newReaderServiceConn(ctx context.Context, cfg *config.Config) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(constants.BACKOFF_LINEAR)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(constants.BACKOFF_RETRIES),
	}

	readerServiceConn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("localhost:%s", cfg.GRPC.ReaderServicePort),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}

	return readerServiceConn, nil
}
