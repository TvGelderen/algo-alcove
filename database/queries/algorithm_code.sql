-- name: CreateAlgorithmCode :exec
INSERT INTO algorithm_code (algorithm_id, language, code)
VALUES ($1, $2, $3);

-- name: UpdateAlgorithmCode :exec
UPDATE algorithm_code
SET language = $2,
    code = $3
WHERE id = $1;

-- name: GetAlgorithmCodeByAgorithmId :many
SELECT * FROM algorithm_code
WHERE algorithm_id = $1;
