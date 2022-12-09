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

func (g *ProductPostgresGateway) CreateProduct(ctx context.Context, p product.Product) error {
	product, err := g.pgRepo.CreateProduct(ctx, &p)
	if err != nil {
		return err
	}

	fmt.Printf("Product %s created", product.Name)

	return nil
}

func (g *ProductPostgresGateway) DeleteProductByID(ctx context.Context, productID string) error {
	if err := g.pgRepo.DeleteProductByID(ctx, productID); err != nil {
		return err
	}

	fmt.Printf("Product with ID: %s has been deleted", productID)

	return nil
}
