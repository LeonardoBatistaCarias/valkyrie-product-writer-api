package product

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product/persistence"
)

type ProductPostgresGateway struct {
	pgRepo persistence.Repository
}

func NewProductPostgresGateway(pgRepo persistence.Repository) *ProductPostgresGateway {
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

func (g *ProductPostgresGateway) DeactivateProductByID(ctx context.Context, productID string) error {
	if err := g.pgRepo.DeleteProductByID(ctx, productID); err != nil {
		return err
	}

	fmt.Printf("Product with ID: %s has been deleted", productID)

	return nil
}

func (g *ProductPostgresGateway) UpdateProductByID(ctx context.Context, p product.Product) error {
	product, err := g.pgRepo.UpdateProductByID(ctx, &p)
	if err != nil {
		return err
	}

	fmt.Printf("Product %s updated", product.Name)

	return nil
}
