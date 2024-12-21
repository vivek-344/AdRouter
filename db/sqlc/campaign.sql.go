// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: campaign.sql

package db

import (
	"context"
)

const createCampaign = `-- name: CreateCampaign :one
INSERT INTO campaign (
  cid,
  name,
  img,
  cta
) VALUES (
    $1, $2, $3, $4
)
RETURNING cid, name, img, cta, status, created_at
`

type CreateCampaignParams struct {
	Cid  string `json:"cid"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Cta  string `json:"cta"`
}

func (q *Queries) CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, createCampaign,
		arg.Cid,
		arg.Name,
		arg.Img,
		arg.Cta,
	)
	var i Campaign
	err := row.Scan(
		&i.Cid,
		&i.Name,
		&i.Img,
		&i.Cta,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCampaign = `-- name: DeleteCampaign :exec
DELETE FROM campaign
WHERE cid = $1
`

func (q *Queries) DeleteCampaign(ctx context.Context, cid string) error {
	_, err := q.db.Exec(ctx, deleteCampaign, cid)
	return err
}

const getCampaign = `-- name: GetCampaign :one
SELECT cid, name, img, cta, status, created_at
FROM campaign
WHERE cid = $1
`

func (q *Queries) GetCampaign(ctx context.Context, cid string) (Campaign, error) {
	row := q.db.QueryRow(ctx, getCampaign, cid)
	var i Campaign
	err := row.Scan(
		&i.Cid,
		&i.Name,
		&i.Img,
		&i.Cta,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listAllActiveCampaign = `-- name: ListAllActiveCampaign :many
SELECT cid, name, img, cta, status, created_at
FROM campaign
WHERE status = 'active'
`

func (q *Queries) ListAllActiveCampaign(ctx context.Context) ([]Campaign, error) {
	rows, err := q.db.Query(ctx, listAllActiveCampaign)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Campaign{}
	for rows.Next() {
		var i Campaign
		if err := rows.Scan(
			&i.Cid,
			&i.Name,
			&i.Img,
			&i.Cta,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllCampaign = `-- name: ListAllCampaign :many
SELECT cid, name, img, cta, status, created_at
FROM campaign
`

func (q *Queries) ListAllCampaign(ctx context.Context) ([]Campaign, error) {
	rows, err := q.db.Query(ctx, listAllCampaign)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Campaign{}
	for rows.Next() {
		var i Campaign
		if err := rows.Scan(
			&i.Cid,
			&i.Name,
			&i.Img,
			&i.Cta,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const toggleStatus = `-- name: toggleStatus :one
UPDATE campaign
SET status = CASE 
    WHEN status = 'active' THEN 'inactive'
    ELSE 'inactive'
END
WHERE cid = $1
RETURNING status
`

func (q *Queries) toggleStatus(ctx context.Context, cid string) (StatusType, error) {
	row := q.db.QueryRow(ctx, toggleStatus, cid)
	var status StatusType
	err := row.Scan(&status)
	return status, err
}

const updateCampaignCta = `-- name: updateCampaignCta :one
UPDATE campaign
SET cta = $2
WHERE cid = $1
RETURNING cid, name, img, cta, status, created_at
`

type updateCampaignCtaParams struct {
	Cid string `json:"cid"`
	Cta string `json:"cta"`
}

func (q *Queries) updateCampaignCta(ctx context.Context, arg updateCampaignCtaParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, updateCampaignCta, arg.Cid, arg.Cta)
	var i Campaign
	err := row.Scan(
		&i.Cid,
		&i.Name,
		&i.Img,
		&i.Cta,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateCampaignImage = `-- name: updateCampaignImage :one
UPDATE campaign
SET img = $2
WHERE cid = $1
RETURNING cid, name, img, cta, status, created_at
`

type updateCampaignImageParams struct {
	Cid string `json:"cid"`
	Img string `json:"img"`
}

func (q *Queries) updateCampaignImage(ctx context.Context, arg updateCampaignImageParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, updateCampaignImage, arg.Cid, arg.Img)
	var i Campaign
	err := row.Scan(
		&i.Cid,
		&i.Name,
		&i.Img,
		&i.Cta,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateCampaignName = `-- name: updateCampaignName :one
UPDATE campaign
SET name = $2
WHERE cid = $1
RETURNING cid, name, img, cta, status, created_at
`

type updateCampaignNameParams struct {
	Cid  string `json:"cid"`
	Name string `json:"name"`
}

func (q *Queries) updateCampaignName(ctx context.Context, arg updateCampaignNameParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, updateCampaignName, arg.Cid, arg.Name)
	var i Campaign
	err := row.Scan(
		&i.Cid,
		&i.Name,
		&i.Img,
		&i.Cta,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
