-- name: GetMeasureByID :one
SELECT * FROM measures
WHERE id = $1 LIMIT 1;

-- name: GetMeasureByCustomerID :many
SELECT * FROM measures
WHERE customer_id = $3
LIMIT $1
OFFSET $2;

-- name: Listmeasures :many
SELECT * FROM measures
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateMeasure :one
INSERT INTO measures (
  customer_id, name, number
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteMeasureByID :exec
DELETE FROM measures
WHERE id = $1;

-- name: DeleteMeasureByCustomerID :exec
DELETE FROM measures
WHERE customer_id = $1;

-- name: UpdateMeasureNumber :one
UPDATE measures SET number = $2
WHERE id = $1
RETURNING *;

