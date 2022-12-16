package product_processor

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/create"
	model "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/model/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/reader_service/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	"github.com/avast/retry-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (p *ProductMessageProcessor) processCreateProduct(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &model.Product{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		fmt.Errorf("proto.Unmarshal", err)
		p.commitErrMessage(ctx, r, m)
		return
	}

	command := create.NewCreateProductCommand(msg.GetName(), msg.GetDescription(), 1, msg.GetPrice(), msg.GetQuantity(), uuid.NewV4(), nil, true)

	if err := retry.Do(func() error {
		return p.commands.CreateProduct.Handle(ctx, *command)
	}, append(message_processor.RETRY_OPTIONS, retry.Context(ctx))...); err != nil {
		fmt.Errorf("CreateProduct.Handle", err)
		return
	}

	connection, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewProductReaderServiceClient(connection)
	res, err := client.CreateProduct(context.Background(), &pb.CreateProductReq{Product: msg})
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	log.Print(res)

	p.commitMessage(ctx, r, m)
}
