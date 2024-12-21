// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: target_country.sql

package db

import (
	"context"
)

const createTargetCountry = `-- name: CreateTargetCountry :one
INSERT INTO target_country (
    cid,
    country,
    rule
) VALUES (
    $1, $2, $3
)
RETURNING cid, country, rule
`

type CreateTargetCountryParams struct {
	Cid     string   `json:"cid"`
	Country string   `json:"country"`
	Rule    RuleType `json:"rule"`
}

func (q *Queries) CreateTargetCountry(ctx context.Context, arg CreateTargetCountryParams) (TargetCountry, error) {
	row := q.db.QueryRow(ctx, createTargetCountry, arg.Cid, arg.Country, arg.Rule)
	var i TargetCountry
	err := row.Scan(&i.Cid, &i.Country, &i.Rule)
	return i, err
}

const deleteTargetCountry = `-- name: DeleteTargetCountry :exec
DELETE FROM target_country
WHERE cid = $1
`

func (q *Queries) DeleteTargetCountry(ctx context.Context, cid string) error {
	_, err := q.db.Exec(ctx, deleteTargetCountry, cid)
	return err
}

const getTargetCountry = `-- name: GetTargetCountry :one
SELECT cid, country, rule
FROM target_country
WHERE cid = $1
`

func (q *Queries) GetTargetCountry(ctx context.Context, cid string) (TargetCountry, error) {
	row := q.db.QueryRow(ctx, getTargetCountry, cid)
	var i TargetCountry
	err := row.Scan(&i.Cid, &i.Country, &i.Rule)
	return i, err
}

const updateTargetCountry = `-- name: updateTargetCountry :one
UPDATE target_country
SET country = $2, rule = $3
WHERE cid = $1
RETURNING cid, country, rule
`

type updateTargetCountryParams struct {
	Cid     string   `json:"cid"`
	Country string   `json:"country"`
	Rule    RuleType `json:"rule"`
}

func (q *Queries) updateTargetCountry(ctx context.Context, arg updateTargetCountryParams) (TargetCountry, error) {
	row := q.db.QueryRow(ctx, updateTargetCountry, arg.Cid, arg.Country, arg.Rule)
	var i TargetCountry
	err := row.Scan(&i.Cid, &i.Country, &i.Rule)
	return i, err
}
