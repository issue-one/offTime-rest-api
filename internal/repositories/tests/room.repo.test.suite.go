package tests

import (
	"context"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	internalModels "github.com/issue-one/offTime-rest-api/internal/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
	. "github.com/issue-one/offTime-rest-api/internal/repositories/mock"
)

func RoomRepositoryTestSuite(
	t *testing.T,
	getRepo func() (repositories.Room, error),
	disposeOfRepo func(repo repositories.Room) error) {
	ctx := context.TODO()

	testRoomname := "Bear Town"

	tests := []struct {
		name string
		test func(t *testing.T, repo repositories.Room)
	}{
		{
			name: "GetRoom - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				room, err := repo.GetRoom(ctx, Room01.ID)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if room.ID != Room01.ID {
					t.Errorf("failed: %v != %v", room.ID, Room01.ID)
				}
				if room.Name != Room01.Name {
					t.Errorf("failed: %v != %v", room.Name, Room01.Name)
				}
				if room.HostUsername != Room01.HostUsername {
					t.Errorf("failed: %v != %v", room.HostUsername, Room01.HostUsername)
				}
				if room.StartTime != Room01.StartTime {
					t.Errorf("failed: %v != %v", room.StartTime, Room01.StartTime)
				}
				if room.EndTime != Room01.EndTime {
					t.Errorf("failed: %v != %v", room.EndTime, Room01.EndTime)
				}
				if !cmp.Equal(room.UserUsages, Room01.UserUsages) {
					t.Errorf("failed: %v != %v", room.UserUsages, Room01.UserUsages)
				}
			},
		},
		{
			name: "GetRoom - it throws ErrRoomNotFound if Room not found",
			test: func(t *testing.T, repo repositories.Room) {
				_, err := repo.GetRoom(ctx, strfmt.UUID(uuid.New().String()))
				if err == nil || err != repositories.ErrRoomNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
		{
			name: "CreateRoom - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				room, err := repo.CreateRoom(ctx, User01.Username, testRoomname)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if room.Name != testRoomname {
					t.Errorf("failed: %v != %v", room.Name, testRoomname)
				}
				if room.HostUsername != User01.Username {
					t.Errorf("failed: %v != %v", room.HostUsername, User01.Username)
				}
			},
		},
		{
			name: "CreateRoom - it assigns CreatedAt",
			test: func(t *testing.T, repo repositories.Room) {
				room, err := repo.CreateRoom(ctx, User01.Username, testRoomname)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if room.CreatedAt.Equal(strfmt.DateTime{}) {
					t.Errorf("failed: cratedAt is nil\n%v", room)
				}
			},
		},
		{
			name: "CreateRoom - it assigns UpdatedAt",
			test: func(t *testing.T, repo repositories.Room) {
				room, err := repo.CreateRoom(ctx, User01.Username, testRoomname)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if room.UpdatedAt.Equal(strfmt.DateTime{}) {
					t.Errorf("failed: cratedAt is nil\n%v", room)
				}
			},
		},
		{
			name: "UpdateRoom - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				update := internalModels.UpdateRoomInput{
					Name:      testRoomname,
					EndTime:   strfmt.DateTime(time.Now().Round(0)),
					StartTime: strfmt.DateTime(time.Now().Round(0)),
				}
				room, err := repo.UpdateRoom(ctx, Room01.ID, &update)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if room.Name != update.Name {
					t.Errorf("failed: %v != %v", room.Name, update.Name)
				}
				if room.EndTime != update.EndTime {
					t.Errorf("failed: %v != %v", room.EndTime, update.EndTime)
				}
			},
		},
		{
			name: "UpdateRoom - it throws ErrRoomNotFound if Room not found",
			test: func(t *testing.T, repo repositories.Room) {
				_, err := repo.UpdateRoom(ctx, strfmt.UUID(uuid.New().String()), &internalModels.UpdateRoomInput{
					Name: testRoomname,
				})
				if err == nil || err != repositories.ErrRoomNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
		{
			name: "GetAllRooms - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				rooms, _, err := repo.GetAllRooms(ctx, 10, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if rooms == nil {
					t.Error("nil array returned")
					return
				}
				if len(rooms) != 3 {
					t.Errorf("only %v were returned", len(rooms))
				}
			},
		},
		{
			name: "GetAllRooms - it returns totalCount",
			test: func(t *testing.T, repo repositories.Room) {
				_, totalCount, err := repo.GetAllRooms(ctx, 10, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if totalCount != 3 {
					t.Errorf("totalCount returned %v != 3", totalCount)
				}
			},
		},
		{
			name: "GetAllRooms - it limits",
			test: func(t *testing.T, repo repositories.Room) {
				rooms, _, err := repo.GetAllRooms(ctx, 2, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(rooms) != 2 {
					t.Errorf("limit wasn't respected: %v were returned", len(rooms))
				}
			},
		},
		{
			name: "GetAllRooms - it offsets",
			test: func(t *testing.T, repo repositories.Room) {
				rooms, _, err := repo.GetAllRooms(ctx, 10, 2)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(rooms) != 1 {
					t.Errorf("offset wasn't respected: %v were returned", len(rooms))
				}
			},
		},
		{
			name: "GetMultipleRooms - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				rooms, err := repo.GetMultipleRooms(ctx, []strfmt.UUID{Room01.ID, Room02.ID})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if rooms == nil {
					t.Error("nil array returned")
					return
				}
				if len(rooms) != 2 {
					t.Errorf("only %v were returned", len(rooms))
				}
			},
		},

		{
			name: "GetMultipleRooms - it throws ErrRoomNotFound if Room not found",
			test: func(t *testing.T, repo repositories.Room) {
				_, err := repo.GetMultipleRooms(ctx, []strfmt.UUID{Room01.ID, strfmt.UUID(uuid.New().String()), Room02.ID})
				if err == nil || err != repositories.ErrRoomNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
		{
			name: "UpdateRoomUserUsages - it succeeds",
			test: func(t *testing.T, repo repositories.Room) {
				update := map[string]int64{
					User01.Username: int64(time.Second * 10),
					User02.Username: int64(time.Second * 100),
					User03.Username: int64(time.Second * 1000),
				}
				room, err := repo.UpdateRoomUserUsages(ctx, Room03.ID, &update)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				for _, usage := range room.UserUsages {
					if usage.TotalTime != update[usage.Username] {
						t.Errorf("failed: usage time's are not equal => %v:%v != %v:%v",
							usage.Username, usage.TotalTime, usage.Username, update[usage.Username])
					}
				}
			},
		},
		{
			name: "UpdateRoomUserUsages - it adds new usage instance if user is new",
			test: func(t *testing.T, repo repositories.Room) {
				update := map[string]int64{
					User02.Username: int64(time.Second * 100),
				}
				room, err := repo.UpdateRoomUserUsages(ctx, Room03.ID, &update)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				for _, usage := range room.UserUsages {
					if usage.Username != User01.Username {
						return
					}
				}
				t.Errorf("failed: entry wasn't addded => %v != %v", room.UserUsages, update)
			},
		},
		{
			name: "UpdateRoomUserUsages - it throws ErrRoomNotFound if Room not found",
			test: func(t *testing.T, repo repositories.Room) {
				_, err := repo.UpdateRoomUserUsages(ctx, strfmt.UUID(uuid.New().String()), &map[string]int64{
					User01.Username: int64(time.Second * 100),
				})
				if err == nil || err != repositories.ErrRoomNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo, err := getRepo()
			if err != nil {
				t.Errorf("error getting repo: %v", err)
			}
			testCase.test(t, repo)
			err = disposeOfRepo(repo)
			if err != nil {
				t.Errorf("error disposing of repo: %v", err)
			}
		})
	}

}
