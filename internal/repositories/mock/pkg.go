package mock

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/issue-one/offTime-rest-api/gen/models"
)

var room1Id = strfmt.UUID(uuid.New().String())
var room2Id = strfmt.UUID(uuid.New().String())
var room3Id = strfmt.UUID(uuid.New().String())

var (
	User01 = models.User{
		Username:   "almaz",
		Email:      "pearl@pax.os",
		Password:   "thejewel",
		PictureURL: "cam_123124.jpg",
		RoomHistory: []strfmt.UUID{
			room1Id, room3Id},
		CreatedAt: strfmt.DateTime(time.Now().Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Round(0)),
	}
	User02 = models.User{
		Username:    "tseahay",
		Email:       "sun@worsh.ip",
		Password:    "kokob",
		PictureURL:  "inyoface.jpg",
		RoomHistory: []strfmt.UUID{room2Id, room3Id},
		CreatedAt:   strfmt.DateTime(time.Now().Round(0)),
		UpdatedAt:   strfmt.DateTime(time.Now().Round(0)),
	}
	User03 = models.User{
		Username:   "emebet",
		Email:      "mistress@nig.ht",
		Password:   "danger",
		PictureURL: "perplex.jpg",
		RoomHistory: []strfmt.UUID{
			room1Id, room2Id, room3Id,
		},
		CreatedAt: strfmt.DateTime(time.Now().Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Round(0)),
	}
)

var (
	Room01 = models.Room{
		ID:           room1Id,
		HostUsername: User01.Username,
		Name:         "The Garden",
		UserUsages: []*models.RoomUserUsagesItems0{
			{
				Username:  User01.Username,
				TotalTime: int64(time.Minute * 3),
			},
			{
				Username:  User03.Username,
				TotalTime: int64(time.Minute * 15),
			},
		},
		StartTime: strfmt.DateTime(time.Now().Add(-time.Hour * 75).Round(0)),
		EndTime:   strfmt.DateTime(time.Now().Add(-time.Hour * 72).Round(0)),
		CreatedAt: strfmt.DateTime(time.Now().Add(-time.Hour * 75).Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Add(-time.Hour * 72).Round(0)),
	}
	Room02 = models.Room{
		ID:           room2Id,
		HostUsername: User02.Username,
		Name:         "HAOB Lounge",
		UserUsages: []*models.RoomUserUsagesItems0{
			{
				Username:  User02.Username,
				TotalTime: int64(time.Second * 45),
			},
			{
				Username:  User03.Username,
				TotalTime: int64(time.Second * 150),
			},
		},
		StartTime: strfmt.DateTime(time.Now().Add(-time.Hour * 3).Round(0)),
		EndTime:   strfmt.DateTime(time.Now().Add(-time.Hour * 1).Round(0)),
		CreatedAt: strfmt.DateTime(time.Now().Add(-time.Hour * 3).Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Add(-time.Hour * 1).Round(0)),
	}
	// Room03 is ongoing
	Room03 = models.Room{
		ID:           room3Id,
		HostUsername: User03.Username,
		Name:         "Northgate Artifactory",
		UserUsages: []*models.RoomUserUsagesItems0{
			{
				Username:  User03.Username,
				TotalTime: int64(time.Minute * 2),
			},
			{
				Username:  User01.Username,
				TotalTime: int64(time.Minute * 7),
			},
			{
				Username:  User02.Username,
				TotalTime: int64(time.Minute * 4),
			},
		},
		StartTime: strfmt.DateTime(time.Now().Add(-time.Minute * 30).Round(0)),
		CreatedAt: strfmt.DateTime(time.Now().Add(-time.Minute * 75).Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Round(0)),
	}
)
