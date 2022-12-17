package product_processor

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/update"
	model "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/model/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/grpc/reader_service/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/kafka/message_processor"
	"github.com/avast/retry-go"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (p *ProductMessageProcessor) processUpdateProductByID(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &model.Product{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		p.log.WarnMsg("proto.Unmarshal", err)
		p.commitErrMessage(ctx, r, m)
		return
	}

	command := update.NewUpdateProductByIDCommand(uuid.FromStringOrNil(msg.GetProductID()), msg.GetName(), msg.GetDescription(), 1, msg.GetPrice(), msg.GetQuantity(), uuid.FromStringOrNil(msg.GetCategoryID()), nil, true)

	if err := retry.Do(func() error {
		return p.commands.UpdateProductByID.Handle(ctx, *command)
	}, append(message_processor.RETRY_OPTIONS, retry.Context(ctx))...); err != nil {
		p.log.WarnMsg("UpdateProductByID.Handle", err)
		return
	}

	res, err := p.rc.UpdateProductByID(context.Background(), &pb.UpdateProductByIDReq{Product: msg})
	if err != nil {
		p.metrics.ErrorGrpcRequests.Inc()
		p.log.Errorf("Error in UpdateProductByID Grpc call to Product Reader Service", err)
	} else {
		p.metrics.SuccessGrpcRequests.Inc()
	}

	p.log.Infof("Product has been updated", res)

	p.commitMessage(ctx, r, m)
}
