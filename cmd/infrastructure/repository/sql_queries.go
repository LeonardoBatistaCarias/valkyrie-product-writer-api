package repository

const (
	createProductQuery = `INSERT INTO products (product_id, name, description, price, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, now(), now()) RETURNING product_id, name, description, price, created_at, updated_at`

	deleteProductByIdQuery = `UPDATE products p SET active=false WHERE product_id=$1`

	updateProductQuery = `UPDATE products p SET 
                      name=COALESCE(NULLIF($1, ''), name), 
                      description=COALESCE(NULLIF($2, ''), description), 
                      price=COALESCE(NULLIF($3, 0), price),
                      updated_at = now()
                      WHERE product_id=$4
                      RETURNING product_id, name, description, price, created_at, updated_at`
)
