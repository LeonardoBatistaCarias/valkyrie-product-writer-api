package persistence

const (
	CREATE_PRODUCT_QUERY = `INSERT INTO products (product_id, name, description, price, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, now(), now()) RETURNING product_id, name, description, price, created_at, updated_at`

	DEACTIVATE_PRODUCT_BY_ID_QUERY = `UPDATE products p SET active=false WHERE product_id=$1`

	DELETE_PRODUCT_BY_ID = `DELETE FROM products WHERE product_id = $1`

	UPDATE_PRODUCT_QUERY = `UPDATE products p SET 
                      name=COALESCE(NULLIF($1, ''), name), 
                      description=COALESCE(NULLIF($2, ''), description), 
                      price=COALESCE(NULLIF($3, 0), price),
                      updated_at = now()
                      WHERE product_id=$4
                      RETURNING product_id, name, description, price, created_at, updated_at`
)
