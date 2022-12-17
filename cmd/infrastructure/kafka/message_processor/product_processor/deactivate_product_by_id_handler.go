package product_processor

import (
	"context"
	"fmt"
	deleteKafkaMessage "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/pb/kafka"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/pb/reader_service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (p *ProductMessageProcessor) processDeactivateProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &deleteKafkaMessage.DeactivateProductByIDKafkaMessage{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		p.log.WarnMsg("proto.Unmarshal", err)
		p.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return p.commands.DeactivateProductByID.Handle(ctx, msg.GetProductID())
	}, append(message_processor.RETRY_OPTIONS, retry.Context(ctx))...); err != nil {
		p.log.WarnMsg("DeactivateProductByID.Handle", err)
		return
	}

	if _, err := p.rc.DeactivateProductByID(context.Background(), &reader_service.DeactivateProductByIDReq{ProductID: msg.GetProductID()}); err != nil {
		p.metrics.ErrorGrpcRequests.Inc()
		p.log.Errorf("Error in DeactivateProductByID Grpc call to Product Reader Service", err)
	} else {
		p.metrics.SuccessGrpcRequests.Inc()
	}

	p.log.Info(fmt.Printf("The Product with ID %s deactivated in Postgres", msg.GetProductID()))

	p.commitMessage(ctx, r, m)
}
