package product

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/product/persistence"
)

type PostgresGateway struct {
	pgRepo persistence.Repository
}

func NewProductPostgresGateway(pgRepo persistence.Repository) *PostgresGateway {
	return &PostgresGateway{
		pgRepo: pgRepo,
	}
}

func (g *PostgresGateway) CreateProduct(ctx context.Context, p *product.Product) error {
	if err := g.pgRepo.Create(ctx, p); err != nil {
		return err
	}
	return nil
}

func (g *PostgresGateway) DeleteProductByID(ctx context.Context, productID string) error {
	if err := g.pgRepo.DeleteByID(ctx, productID); err != nil {
		return err
	}
	return nil
}

func (g *PostgresGateway) DeactivateProductByID(ctx context.Context, productID string) error {
	if err := g.pgRepo.DeactivateByID(ctx, productID); err != nil {
		return err
	}
	return nil
}

func (g *PostgresGateway) UpdateProductByID(ctx context.Context, p *product.Product) error {
	if err := g.pgRepo.UpdateByID(ctx, p); err != nil {
		return err
	}
	return nil
}
