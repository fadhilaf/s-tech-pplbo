-- name: GetUser :many
SELECT id, name, email FROM users
ORDER BY name;

-- name: GetUserByEmail :one
SELECT * FROM users
where email = ? LIMIT 1;


-- name: CreateUser :execresult
INSERT INTO users (
  name, email, password_hash, address, phone
) VALUES ( ?, ?, ?, ?, ? );

-- name: UpdateUser :execresult
UPDATE users SET
  name = ?,
  address = ?,
  phone = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
