package repository

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type Repository interface {
	CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error)
}
