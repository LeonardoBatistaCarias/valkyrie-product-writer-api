package repository

import (
	"context"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type productRepository struct {
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewProductRepository(cfg *config.Config, db *pgxpool.Pool) *productRepository {
	return &productRepository{cfg: cfg, db: db}
}

func (pr *productRepository) CreateProduct(ctx context.Context, p *product.Product) (*product.Product, error) {
	var created product.Product

	if err := pr.db.QueryRow(ctx, createProductQuery, &p.ProductID, &p.Name, &p.Description, &p.Price).Scan(
		&created.ProductID,
		&created.Name,
		&created.Description,
		&created.Price,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("db.QueryRow", err)
	}

	return &created, nil
}

func (pr *productRepository) DeleteProductByID(ctx context.Context, productID string) error {
	if _, err := pr.db.Exec(ctx, deleteProductByIdQuery, productID); err != nil {
		return fmt.Errorf("db.QueryRow %v", err)
	}

	return nil
}

func (pr *productRepository) UpdateProductByID(ctx context.Context, p *product.Product) (*product.Product, error) {
	var updated product.Product

	if err := pr.db.QueryRow(ctx, updateProductQuery, &p.Name, &p.Description, &p.Price, &p.ProductID).Scan(
		&updated.ProductID,
		&updated.Name,
		&updated.Description,
		&updated.Price,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &updated, nil
}
