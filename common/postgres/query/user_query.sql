-- name: GetUser :one
SELECT * FROM users
where id = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :execresult
INSERT INTO users (
  name, email, password_hash
) VALUES ( ?, ?, ? );

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
