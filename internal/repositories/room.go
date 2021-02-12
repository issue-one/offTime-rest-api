package repositories

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	internalModels "github.com/issue-one/offTime-rest-api/internal/models"
)

type Room interface {
	GetAllRooms(ctx context.Context, limit int64, offset int64) (items []*models.Room, totalCount int, err error)
	GetMultipleRooms(ctx context.Context, roomIds []strfmt.UUID) (items []*models.Room, err error)
	CreateRoom(ctx context.Context, username string, name string) (*models.Room, error)
	GetRoom(ctx context.Context, roomid strfmt.UUID) (*models.Room, error)
	UpdateRoom(ctx context.Context, roomid strfmt.UUID, u *internalModels.UpdateRoomInput) (*models.Room, error)
	UpdateRoomUserUsages(ctx context.Context, roomid strfmt.UUID, usages *map[string]int64) (*models.Room, error)
	DeleteRoom(ctx context.Context, roomid strfmt.UUID) error
}
