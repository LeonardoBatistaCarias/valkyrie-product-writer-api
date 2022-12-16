package delete

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type DeleteProductByIDCommandHandler interface {
	Handle(ctx context.Context, productID string) error
}

type deleteProductByIDHandler struct {
	gateway product.ProductGateway
}

func NewDeleteProductByIDHandler(productGateway product.ProductGateway) *deleteProductByIDHandler {
	return &deleteProductByIDHandler{gateway: productGateway}
}

func (c *deleteProductByIDHandler) Handle(ctx context.Context, productID string) error {
	c.gateway.DeleteProductByID(ctx, productID)

	return nil
}
