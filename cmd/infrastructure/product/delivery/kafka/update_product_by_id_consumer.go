package kafka

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/update"
	protoProduct "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/proto/product"
	"github.com/avast/retry-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *productMessageProcessor) processUpdateProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &protoProduct.Product{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		fmt.Errorf("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	command := update.NewUpdateProductByIDCommand(uuid.FromStringOrNil(msg.GetProductID()), msg.GetName(), msg.GetDescription(), 1, msg.GetPrice(), msg.GetQuantity(), uuid.FromStringOrNil(msg.GetCategoryID()), nil, true)

	if err := retry.Do(func() error {
		return s.commands.UpdateProductByID.Handle(ctx, *command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		fmt.Errorf("UpdateProduct.Handle", err)
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
