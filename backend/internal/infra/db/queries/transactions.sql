-- name: CreateTransaction :one
INSERT INTO transactions (
  id,
  description,
  date,
  amount
) VALUES (
  $1, $2, $3, $4
) 
RETURNING *;

-- name: GetTransaction :one
SELECT
  id,
  description,
  date,
  amount
FROM 
    transactions
WHERE 
    id = $1;

-- name: GetLatestTransactions :many
SELECT
    id,
    description,
    date,
    amount
FROM 
    transactions
ORDER BY 
    ctid DESC
LIMIT $1;