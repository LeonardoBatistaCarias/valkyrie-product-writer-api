package persistence

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type Repository interface {
	Create(ctx context.Context, product *product.Product) error
	DeleteByID(ctx context.Context, productID string) error
	DeactivateByID(ctx context.Context, productID string) error
	UpdateByID(ctx context.Context, product *product.Product) error
}
