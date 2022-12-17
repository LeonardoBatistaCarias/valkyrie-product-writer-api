package product_processor

import (
	"context"
	"fmt"
	deleteKafkaMessage "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/kafka/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/reader_service/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (p *ProductMessageProcessor) processDeleteProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &deleteKafkaMessage.DeleteProductByIDKafkaMessage{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		p.log.WarnMsg("proto.Unmarshal", err)
		p.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return p.commands.DeleteProductByID.Handle(ctx, msg.GetProductID())
	}, append(message_processor.RETRY_OPTIONS, retry.Context(ctx))...); err != nil {
		p.log.WarnMsg("DeleteProductByID.Handle", err)
		return
	}

	if _, err := p.rc.DeleteProductByID(context.Background(), &pb.DeleteProductByIDReq{ProductID: msg.GetProductID()}); err != nil {
		p.metrics.ErrorGrpcRequests.Inc()
		p.log.Errorf("Error in DeleteProductByID Grpc call to Product Reader Service", err)
	} else {
		p.metrics.SuccessGrpcRequests.Inc()
	}

	p.log.Info(fmt.Printf("The Product with ID %s has been deleted in Postgres", msg.GetProductID()))

	p.commitMessage(ctx, r, m)
}
