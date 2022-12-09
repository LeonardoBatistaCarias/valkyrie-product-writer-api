package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/create"
	deleteByID "github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/application/commands/delete"
	"github.com/LeonardoBatistaCarias/valkyrie-product-writer-api/cmd/domain/product"
)

type ProductCommands struct {
	CreateProduct     create.CreateProductCommandHandler
	DeleteProductByID deleteByID.DeleteProductByIDCommandHandler
}

func NewProductCommands(pgGateway product.ProductGateway) *ProductCommands {
	createHandler := create.NewCreateProductHandler(pgGateway)
	deleteProductByIDHandler := deleteByID.NewDeleteProductByIDHandler(pgGateway)
	return &ProductCommands{CreateProduct: createHandler, DeleteProductByID: deleteProductByIDHandler}
}
