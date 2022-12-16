package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/create"
	deleteBy "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/delete"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/update"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type ProductCommands struct {
	CreateProduct         create.CreateProductCommandHandler
	DeleteProductByID     deleteBy.DeleteProductByIDCommandHandler
	DeactivateProductByID deleteBy.DeactivateProductByIDCommandHandler
	UpdateProductByID     update.UpdateProductByIDCommandHandler
}

func NewProductCommands(pgGateway product.ProductGateway) *ProductCommands {
	createHandler := create.NewCreateProductHandler(pgGateway)
	deleteProductByIDHandler := deleteBy.NewDeleteProductByIDHandler(pgGateway)
	deactivateProductByIDHandler := deleteBy.NewDeleteProductByIDHandler(pgGateway)
	updateProductByIDHandler := update.NewUpdateProductByIDHandler(pgGateway)

	return &ProductCommands{
		CreateProduct:         createHandler,
		DeleteProductByID:     deleteProductByIDHandler,
		DeactivateProductByID: deactivateProductByIDHandler,
		UpdateProductByID:     updateProductByIDHandler,
	}
}
