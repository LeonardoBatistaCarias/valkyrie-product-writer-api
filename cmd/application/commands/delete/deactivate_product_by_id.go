package delete

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type DeactivateProductByIDCommandHandler interface {
	Handle(ctx context.Context, productID string) error
}

type deactivateProductByIDHandler struct {
	gateway product.ProductGateway
}

func NewDeactivateProductByIDHandler(productGateway product.ProductGateway) *deactivateProductByIDHandler {
	return &deactivateProductByIDHandler{gateway: productGateway}
}

func (c *deactivateProductByIDHandler) Handle(ctx context.Context, productID string) error {
	if err := c.gateway.DeleteProductByID(ctx, productID); err != nil {
		return err
	}

	return nil
}
