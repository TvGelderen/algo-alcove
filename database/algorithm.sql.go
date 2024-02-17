// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: algorithm.sql

package database

import (
	"context"
)

const createAlgorithm = `-- name: CreateAlgorithm :one
INSERT INTO algorithms (text_id, name, type, position, explanation, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, timezone('utc', NOW()), timezone('utc', NOW()))
RETURNING id
`

type CreateAlgorithmParams struct {
	TextID      string
	Name        string
	Type        int16
	Position    int16
	Explanation string
}

func (q *Queries) CreateAlgorithm(ctx context.Context, arg CreateAlgorithmParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createAlgorithm,
		arg.TextID,
		arg.Name,
		arg.Type,
		arg.Position,
		arg.Explanation,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getAlgorithmById = `-- name: GetAlgorithmById :one
SELECT id, text_id, name, type, position, explanation, created_at, updated_at FROM algorithms
WHERE id = $1
`

func (q *Queries) GetAlgorithmById(ctx context.Context, id int32) (Algorithm, error) {
	row := q.db.QueryRowContext(ctx, getAlgorithmById, id)
	var i Algorithm
	err := row.Scan(
		&i.ID,
		&i.TextID,
		&i.Name,
		&i.Type,
		&i.Position,
		&i.Explanation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAlgorithmByTextId = `-- name: GetAlgorithmByTextId :one
SELECT id, text_id, name, type, position, explanation, created_at, updated_at FROM algorithms
WHERE text_id = $1
`

func (q *Queries) GetAlgorithmByTextId(ctx context.Context, textID string) (Algorithm, error) {
	row := q.db.QueryRowContext(ctx, getAlgorithmByTextId, textID)
	var i Algorithm
	err := row.Scan(
		&i.ID,
		&i.TextID,
		&i.Name,
		&i.Type,
		&i.Position,
		&i.Explanation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAlgorithmByType = `-- name: GetAlgorithmByType :one
SELECT id, text_id, name, type, position, explanation, created_at, updated_at FROM algorithms
WHERE type = $1
`

func (q *Queries) GetAlgorithmByType(ctx context.Context, type_ int16) (Algorithm, error) {
	row := q.db.QueryRowContext(ctx, getAlgorithmByType, type_)
	var i Algorithm
	err := row.Scan(
		&i.ID,
		&i.TextID,
		&i.Name,
		&i.Type,
		&i.Position,
		&i.Explanation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAlgorithmNames = `-- name: GetAlgorithmNames :many
SELECT text_id, name, type FROM algorithms
`

type GetAlgorithmNamesRow struct {
	TextID string
	Name   string
	Type   int16
}

func (q *Queries) GetAlgorithmNames(ctx context.Context) ([]GetAlgorithmNamesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAlgorithmNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAlgorithmNamesRow
	for rows.Next() {
		var i GetAlgorithmNamesRow
		if err := rows.Scan(&i.TextID, &i.Name, &i.Type); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAlgorithm = `-- name: UpdateAlgorithm :exec
UPDATE algorithms
SET text_id = $2, 
    name = $3,
    updated_at = timezone('utc', NOW())
WHERE id = $1
`

type UpdateAlgorithmParams struct {
	ID     int32
	TextID string
	Name   string
}

func (q *Queries) UpdateAlgorithm(ctx context.Context, arg UpdateAlgorithmParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithm, arg.ID, arg.TextID, arg.Name)
	return err
}

const updateAlgorithmExplanation = `-- name: UpdateAlgorithmExplanation :exec
UPDATE algorithms
SET explanation = $2,
    updated_at = timezone('utc', NOW())
WHERE id = $1
`

type UpdateAlgorithmExplanationParams struct {
	ID          int32
	Explanation string
}

func (q *Queries) UpdateAlgorithmExplanation(ctx context.Context, arg UpdateAlgorithmExplanationParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithmExplanation, arg.ID, arg.Explanation)
	return err
}

const updateAlgorithmName = `-- name: UpdateAlgorithmName :exec
UPDATE algorithms
SET name = $2, 
    updated_at = timezone('utc', NOW())
WHERE id = $1
`

type UpdateAlgorithmNameParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateAlgorithmName(ctx context.Context, arg UpdateAlgorithmNameParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithmName, arg.ID, arg.Name)
	return err
}

const updateAlgorithmPosition = `-- name: UpdateAlgorithmPosition :exec
UPDATE algorithms
SET position = $2,
    updated_at = timezone('utc', NOW()) 
WHERE id = $1
`

type UpdateAlgorithmPositionParams struct {
	ID       int32
	Position int16
}

func (q *Queries) UpdateAlgorithmPosition(ctx context.Context, arg UpdateAlgorithmPositionParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithmPosition, arg.ID, arg.Position)
	return err
}
