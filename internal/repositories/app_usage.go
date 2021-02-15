package repositories

import (
	"context"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

type AppUsage interface {
	GetAllAppUsages(ctx context.Context, username string, limit, offset int64) (items []*models.AppUsage, totalCount int, err error)
	CreateAppUsage(ctx context.Context, username string, u *models.CreateAppUsageInput) (*models.AppUsage, error)
	DeleteAppUsages(ctx context.Context, username string) error
}
