-- name: SaveDocument :one
INSERT INTO documents (
    name,
    data_format,
    data
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;