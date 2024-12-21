// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: campaign_history.sql

package db

import (
	"context"
)

const createCampaignHistory = `-- name: CreateCampaignHistory :exec
INSERT INTO campaign_history (
    cid,
    field_changed,
    old_value,
    new_value
) VALUES (
    $1, $2, $3, $4
)
`

type CreateCampaignHistoryParams struct {
	Cid          string `json:"cid"`
	FieldChanged string `json:"field_changed"`
	OldValue     string `json:"old_value"`
	NewValue     string `json:"new_value"`
}

func (q *Queries) CreateCampaignHistory(ctx context.Context, arg CreateCampaignHistoryParams) error {
	_, err := q.db.Exec(ctx, createCampaignHistory,
		arg.Cid,
		arg.FieldChanged,
		arg.OldValue,
		arg.NewValue,
	)
	return err
}
