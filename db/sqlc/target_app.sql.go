// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: target_app.sql

package db

import (
	"context"
)

const createTargetApp = `-- name: CreateTargetApp :one
INSERT INTO target_app (
    cid,
    app_id,
    rule
) VALUES (
    $1, $2, $3
)
RETURNING cid, app_id, rule
`

type CreateTargetAppParams struct {
	Cid   string   `json:"cid"`
	AppID string   `json:"app_id"`
	Rule  RuleType `json:"rule"`
}

func (q *Queries) CreateTargetApp(ctx context.Context, arg CreateTargetAppParams) (TargetApp, error) {
	row := q.db.QueryRow(ctx, createTargetApp, arg.Cid, arg.AppID, arg.Rule)
	var i TargetApp
	err := row.Scan(&i.Cid, &i.AppID, &i.Rule)
	return i, err
}

const deleteTargetApp = `-- name: DeleteTargetApp :exec
DELETE FROM target_app
WHERE cid = $1
`

func (q *Queries) DeleteTargetApp(ctx context.Context, cid string) error {
	_, err := q.db.Exec(ctx, deleteTargetApp, cid)
	return err
}

const getTargetApp = `-- name: GetTargetApp :one
SELECT cid, app_id, rule
FROM target_app
WHERE cid = $1
`

func (q *Queries) GetTargetApp(ctx context.Context, cid string) (TargetApp, error) {
	row := q.db.QueryRow(ctx, getTargetApp, cid)
	var i TargetApp
	err := row.Scan(&i.Cid, &i.AppID, &i.Rule)
	return i, err
}

const updateTargetApp = `-- name: updateTargetApp :one
UPDATE target_app
SET app_id = $2, rule = $3
WHERE cid = $1
RETURNING cid, app_id, rule
`

type updateTargetAppParams struct {
	Cid   string   `json:"cid"`
	AppID string   `json:"app_id"`
	Rule  RuleType `json:"rule"`
}

func (q *Queries) updateTargetApp(ctx context.Context, arg updateTargetAppParams) (TargetApp, error) {
	row := q.db.QueryRow(ctx, updateTargetApp, arg.Cid, arg.AppID, arg.Rule)
	var i TargetApp
	err := row.Scan(&i.Cid, &i.AppID, &i.Rule)
	return i, err
}
