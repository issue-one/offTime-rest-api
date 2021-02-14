package mock

import (
	"context"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	internalModels "github.com/issue-one/offTime-rest-api/internal/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
	uuid "github.com/satori/go.uuid"
)

type roomRepo struct {
	mutex *sync.RWMutex
	rooms map[strfmt.UUID]*models.Room
}

func NewMockRoomRepository() repositories.Room {
	return roomRepo{
		mutex: &sync.RWMutex{},
		rooms: map[strfmt.UUID]*models.Room{
			Room01.ID: &Room01,
			Room02.ID: &Room02,
			Room03.ID: &Room03,
		},
	}
}

func NewMockRoomRepositoryCopyEntities() repositories.Room {
	room01Copy := Room01
	room02Copy := Room02
	room03Copy := Room03
	return roomRepo{
		mutex: &sync.RWMutex{},
		rooms: map[strfmt.UUID]*models.Room{
			Room01.ID: &room01Copy,
			Room02.ID: &room02Copy,
			Room03.ID: &room03Copy,
		},
	}
}

func NewMockRoomRepositoryEmpty() repositories.Room {
	return roomRepo{
		mutex: &sync.RWMutex{},
		rooms: make(map[strfmt.UUID]*models.Room),
	}
}
func (r roomRepo) CreateRoom(ctx context.Context, username string, name string) (*models.Room, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	createdAt := time.Now().Round(0)
	newRoom := models.Room{
		Name:         name,
		HostUsername: username,
		ID:           strfmt.UUID(uuid.NewV4().String()),
		CreatedAt:    strfmt.DateTime(createdAt),
		UpdatedAt:    strfmt.DateTime(createdAt),
	}
	r.rooms[newRoom.ID] = &newRoom
	copy := newRoom
	return &copy, nil
}

func (r roomRepo) GetRoom(ctx context.Context, roomid strfmt.UUID) (*models.Room, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	room, ok := r.rooms[roomid]
	if !ok {
		return nil, repositories.ErrRoomNotFound
	}
	copy := *room
	return &copy, nil
}

func (r roomRepo) GetAllRooms(ctx context.Context, limit int64, offset int64) (items []*models.Room, totalCount int, err error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	capacity := len(r.rooms)
	if offset < int64(capacity) {
		capacity = int(offset)
	}
	rooms := make([]*models.Room, 0, limit)
	lastIndex := limit + offset
	var ii int64 = -1
	for _, room := range r.rooms {
		ii++
		if ii < offset {
			continue
		}
		if ii >= lastIndex {
			break
		}

		copy := *room
		rooms = append(rooms, &copy)
	}
	return rooms, len(r.rooms), nil
}

func (r roomRepo) GetMultipleRooms(ctx context.Context, roomIds []strfmt.UUID) (items []*models.Room, err error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	rooms := make([]*models.Room, 0, len(roomIds))
	for _, roomId := range roomIds {
		room, ok := r.rooms[roomId]
		if !ok {
			return nil, repositories.ErrRoomNotFound
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r roomRepo) UpdateRoom(ctx context.Context, roomid strfmt.UUID, u *internalModels.UpdateRoomInput) (*models.Room, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	oldRoom, ok := r.rooms[roomid]
	if !ok {
		return nil, repositories.ErrRoomNotFound
	}
	if u.Name != "" {
		oldRoom.Name = u.Name
	}
	if !u.EndTime.Equal(strfmt.DateTime{}) {
		oldRoom.EndTime = u.EndTime
	}
	if !u.StartTime.Equal(strfmt.DateTime{}) {
		oldRoom.StartTime = u.StartTime
	}
	copy := *oldRoom
	return &copy, nil
}

func (r roomRepo) UpdateRoomUserUsages(ctx context.Context, roomid strfmt.UUID, usages *map[string]int64) (*models.Room, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	oldRoom, ok := r.rooms[roomid]
	if !ok {
		return nil, repositories.ErrRoomNotFound
	}
	for username, seconds := range *usages {
		oldRoom.UserUsages[username] = seconds
	}
	copy := *oldRoom
	return &copy, nil
}

func (r roomRepo) DeleteRoom(ctx context.Context, roomid strfmt.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.rooms, roomid)
	return nil
}
