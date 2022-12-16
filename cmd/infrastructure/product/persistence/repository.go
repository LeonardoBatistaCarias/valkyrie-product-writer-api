package persistence

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type Repository interface {
	CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error)
	DeleteProductByID(ctx context.Context, productID string) error
	DeactivateProductByID(ctx context.Context, productID string) error
	UpdateProductByID(ctx context.Context, product *product.Product) (*product.Product, error)
}
