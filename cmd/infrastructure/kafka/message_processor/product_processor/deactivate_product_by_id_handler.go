package product_processor

import (
	"context"
	"fmt"
	deleteKafkaMessage "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/kafka/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/reader_service/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	"github.com/avast/retry-go"
	"google.golang.org/grpc"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (p *ProductMessageProcessor) processDeactivateProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &deleteKafkaMessage.DeleteProductByIDKafkaMessage{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		fmt.Errorf("proto.Unmarshal", err)
		p.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return p.commands.DeleteProductByID.Handle(ctx, msg.GetProductID())
	}, append(message_processor.RETRY_OPTIONS, retry.Context(ctx))...); err != nil {
		fmt.Errorf("DeactivateProduct.Handle", err)
		return
	}

	connection, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewProductReaderServiceClient(connection)
	res, err := client.DeleteProductByID(context.Background(), &pb.DeleteProductByIDReq{ProductID: msg.GetProductID()})
	if err != nil {
		fmt.Errorf("err: %v", err)
	}

	log.Print(res)

	p.commitMessage(ctx, r, m)
}
