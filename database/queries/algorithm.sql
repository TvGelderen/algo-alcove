-- name: CreateAlgorithm :exec
INSERT INTO algorithms (text_id, name, explanation, created_at, updated_at)
VALUES ($1, $2, $3, timezone('utc', NOW()), timezone('utc', NOW()));

-- name: UpdateAlgorithm :exec
UPDATE algorithms
SET text_id = $2, 
    name = $3,
    updated_at = timezone('utc', NOW()) 
WHERE id = $1;

-- name: UpdateAlgorithmName :exec
UPDATE algorithms
SET name = $2, 
    updated_at = timezone('utc', NOW()) 
WHERE id = $1;

-- name: UpdateAlgorithmExplanation :exec
UPDATE algorithms
SET explanation = $2,
    updated_at = timezone('utc', NOW()) 
WHERE id = $1;

-- name: GetAlgorithmById :one
SELECT * FROM algorithms
WHERE id = $1;
