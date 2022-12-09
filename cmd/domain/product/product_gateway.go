package product

import (
	"context"
)

type ProductGateway interface {
	CreateProduct(ctx context.Context, product Product) error
	DeleteProductByID(ctx context.Context, productID string) error
}
