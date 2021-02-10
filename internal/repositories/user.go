package repositories

import (
	"context"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

type User interface {
	GetAllUsers(ctx context.Context, limit int64, offset int64) (items []*models.User, totalCount int, err error)
	CreateUser(ctx context.Context, username string, u *models.CreateUserInput) (*models.User, error)
	GetUser(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, username string, u *models.UpdateUserInput) (*models.User, error)
	SetImage(ctx context.Context, username string, imageName string) (*models.User, error)
	DeleteUser(ctx context.Context, username string) error
	IsUsernameOccupied(ctx context.Context, username string) (bool, error)
	IsEmailOccupied(ctx context.Context, email string, excludedUsername string) (bool, error)
}
