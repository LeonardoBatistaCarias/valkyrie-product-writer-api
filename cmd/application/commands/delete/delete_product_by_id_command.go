package delete

type DeleteProductByIDCommand struct {
	ProductID string
	Active    bool
}

func NewDeleteProductByIDCommand(
	productID string,
	active bool) *DeleteProductByIDCommand {
	return &DeleteProductByIDCommand{
		ProductID: productID,
		Active:    active,
	}
}
