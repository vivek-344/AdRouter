// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error)
	CreateTargetApp(ctx context.Context, arg CreateTargetAppParams) (TargetApp, error)
	CreateTargetCountry(ctx context.Context, arg CreateTargetCountryParams) (TargetCountry, error)
	CreateTargetOs(ctx context.Context, arg CreateTargetOsParams) (TargetO, error)
	DeleteCampaign(ctx context.Context, cid string) error
	DeleteTargetApp(ctx context.Context, cid string) error
	DeleteTargetCountry(ctx context.Context, cid string) error
	DeleteTargetOs(ctx context.Context, cid string) error
	GetCampaign(ctx context.Context, cid string) (Campaign, error)
	GetTargetApp(ctx context.Context, cid string) (TargetApp, error)
	GetTargetCountry(ctx context.Context, cid string) (TargetCountry, error)
	GetTargetOs(ctx context.Context, cid string) (TargetO, error)
	ListAllActiveCampaign(ctx context.Context) ([]Campaign, error)
	ListAllCampaign(ctx context.Context) ([]Campaign, error)
	createCampaignHistory(ctx context.Context, arg createCampaignHistoryParams) error
	toggleStatus(ctx context.Context, cid string) (StatusType, error)
	updateCampaignCta(ctx context.Context, arg updateCampaignCtaParams) (Campaign, error)
	updateCampaignImage(ctx context.Context, arg updateCampaignImageParams) (Campaign, error)
	updateCampaignName(ctx context.Context, arg updateCampaignNameParams) (Campaign, error)
	updateTargetApp(ctx context.Context, arg updateTargetAppParams) (TargetApp, error)
	updateTargetCountry(ctx context.Context, arg updateTargetCountryParams) (TargetCountry, error)
	updateTargetOs(ctx context.Context, arg updateTargetOsParams) (TargetO, error)
}

var _ Querier = (*Queries)(nil)
