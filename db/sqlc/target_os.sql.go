// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: target_os.sql

package db

import (
	"context"
)

const createTargetOs = `-- name: CreateTargetOs :one
INSERT INTO target_os (
    cid,
    os,
    rule
) VALUES (
    $1, $2, $3
)
RETURNING cid, os, rule
`

type CreateTargetOsParams struct {
	Cid  string   `json:"cid"`
	Os   string   `json:"os"`
	Rule RuleType `json:"rule"`
}

func (q *Queries) CreateTargetOs(ctx context.Context, arg CreateTargetOsParams) (TargetOs, error) {
	row := q.db.QueryRow(ctx, createTargetOs, arg.Cid, arg.Os, arg.Rule)
	var i TargetOs
	err := row.Scan(&i.Cid, &i.Os, &i.Rule)
	return i, err
}

const deleteTargetOs = `-- name: DeleteTargetOs :exec
DELETE FROM target_os
WHERE cid = $1
`

func (q *Queries) DeleteTargetOs(ctx context.Context, cid string) error {
	_, err := q.db.Exec(ctx, deleteTargetOs, cid)
	return err
}

const getTargetOs = `-- name: GetTargetOs :one
SELECT cid, os, rule
FROM target_os
WHERE cid = $1
`

func (q *Queries) GetTargetOs(ctx context.Context, cid string) (TargetOs, error) {
	row := q.db.QueryRow(ctx, getTargetOs, cid)
	var i TargetOs
	err := row.Scan(&i.Cid, &i.Os, &i.Rule)
	return i, err
}

const updateTargetOs = `-- name: updateTargetOs :one
UPDATE target_os
SET os = $2, rule = $3
WHERE cid = $1
RETURNING cid, os, rule
`

type updateTargetOsParams struct {
	Cid  string   `json:"cid"`
	Os   string   `json:"os"`
	Rule RuleType `json:"rule"`
}

func (q *Queries) updateTargetOs(ctx context.Context, arg updateTargetOsParams) (TargetOs, error) {
	row := q.db.QueryRow(ctx, updateTargetOs, arg.Cid, arg.Os, arg.Rule)
	var i TargetOs
	err := row.Scan(&i.Cid, &i.Os, &i.Rule)
	return i, err
}
