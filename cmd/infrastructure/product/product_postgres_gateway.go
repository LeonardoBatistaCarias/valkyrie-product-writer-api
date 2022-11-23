package product

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/repository"
)

type ProductPostgresGateway struct {
	pgRepo repository.Repository
}

func NewProductPostgresGateway(pgRepo repository.Repository) *ProductPostgresGateway {
	return &ProductPostgresGateway{
		pgRepo: pgRepo,
	}
}

func (g *ProductPostgresGateway) Create(ctx context.Context, p product.Product) error {
	product, err := g.pgRepo.CreateProduct(ctx, &p)
	if err != nil {
		return err
	}

	fmt.Println("Product created %s", product.Name)

	return nil
}
