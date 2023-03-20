-- name: CreateProduct :execresult
INSERT INTO products (
  name, price, stock, is_service, description, image_url
) VALUES (
  $1, $2, $3, $4, $5, $6
);

-- name: GetProduct :many
SELECT * FROM products
ORDER BY name;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductByQuery :many
SELECT * FROM products
WHERE name ILIKE $1

-- name: UpdateProduct :execresult
UPDATE products SET
  name = $2,
  price = $3,
  stock = $4,
  description = $6,
  image_url = $7
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
