package persistence

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/infrastructure/utils/persistence"
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

func (pr *productRepository) Create(ctx context.Context, p *product.Product) error {
	var created product.Product

	if err := pr.db.QueryRow(ctx, persistence.CREATE_PRODUCT_QUERY, &p.ProductID, &p.Name, &p.Description, &p.Price).Scan(
		&created.ProductID,
		&created.Name,
		&created.Description,
		&created.Price,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return errors.Wrap(err, "db.QueryRow")
	}

	return nil
}

func (pr *productRepository) DeleteByID(ctx context.Context, productID string) error {
	if _, err := pr.db.Exec(ctx, persistence.DELETE_PRODUCT_BY_ID, productID); err != nil {
		return errors.Wrap(err, "db.Exec")
	}

	return nil
}

func (pr *productRepository) DeactivateByID(ctx context.Context, productID string) error {
	if _, err := pr.db.Exec(ctx, persistence.DEACTIVATE_PRODUCT_BY_ID_QUERY, productID); err != nil {
		return errors.Wrap(err, "db.Exec")
	}

	return nil
}

func (pr *productRepository) UpdateByID(ctx context.Context, p *product.Product) error {
	var updated product.Product

	if err := pr.db.QueryRow(ctx, persistence.UPDATE_PRODUCT_QUERY, &p.Name, &p.Description, &p.Price, &p.ProductID).Scan(
		&updated.ProductID,
		&updated.Name,
		&updated.Description,
		&updated.Price,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	); err != nil {
		return errors.Wrap(err, "db.QueryRow")
	}

	return nil
}
