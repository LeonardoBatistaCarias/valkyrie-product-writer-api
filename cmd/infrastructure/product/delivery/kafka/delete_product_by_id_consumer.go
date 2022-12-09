package kafka

import (
	"context"
	"fmt"
	deleteByID "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/delete"
	protoProduct "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/proto/product"
	"github.com/avast/retry-go"
	"google.golang.org/grpc"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *productMessageProcessor) processDeleteProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &protoProduct.Product{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		fmt.Errorf("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	command := deleteByID.NewDeleteProductByIDCommand(msg.GetProductID(), msg.GetActive())

	if err := retry.Do(func() error {
		return s.commands.DeleteProductByID.Handle(ctx, *command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		fmt.Errorf("DeleteProduct.Handle", err)
		return
	}

	connection, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := protoProduct.NewProductServiceClient(connection)
	res, err := client.CreateProduct(context.Background(), msg)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	log.Print(res)

	s.commitMessage(ctx, r, m)
}
