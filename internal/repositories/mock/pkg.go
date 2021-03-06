package mock

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	uuid "github.com/satori/go.uuid"
)

var room1Id = strfmt.UUID(uuid.NewV4().String())
var room2Id = strfmt.UUID(uuid.NewV4().String())
var room3Id = strfmt.UUID(uuid.NewV4().String())

var (
	User01 = models.User{
		Username:   "almaz",
		Email:      "pearl@pax.os",
		PictureURL: "cam_123124.jpg",
		RoomHistory: []strfmt.UUID{
			room1Id, room3Id},
		CreatedAt: strfmt.DateTime(time.Now().Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Round(0)),
	}
	User02 = models.User{
		Username:    "tsehay",
		Email:       "sun@worsh.ip",
		PictureURL:  "inyoface.jpg",
		RoomHistory: []strfmt.UUID{room2Id, room3Id},
		CreatedAt:   strfmt.DateTime(time.Now().Round(0)),
		UpdatedAt:   strfmt.DateTime(time.Now().Round(0)),
	}
	User03 = models.User{
		Username:   "emebet",
		Email:      "mistress@nig.ht",
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
		UserUsages: map[string]int64{
			User01.Username: int64(time.Minute * 3),
			User03.Username: int64(time.Minute * 15),
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
		UserUsages: map[string]int64{
			User02.Username: int64(time.Second * 45),
			User03.Username: int64(time.Second * 150),
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
		UserUsages: map[string]int64{
			User03.Username: int64(time.Minute * 2),
			User01.Username: int64(time.Minute * 7),
			User02.Username: int64(time.Minute * 4),
		},
		StartTime: strfmt.DateTime(time.Now().Add(-time.Minute * 30).Round(0)),
		CreatedAt: strfmt.DateTime(time.Now().Add(-time.Minute * 75).Round(0)),
		UpdatedAt: strfmt.DateTime(time.Now().Round(0)),
	}
)

var (
	AppUsage01 = models.AppUsage{
		AppName:        "Telegram",
		AppPackageName: "com.telegram.messenger",
		CreatedAt:      strfmt.DateTime(time.Now().Round(0)),
		DateOfUse:      strfmt.Date(time.Now().Round(0)),
		TimeDuration:   int64(time.Minute * 45),
	}
	AppUsage02 = models.AppUsage{
		AppName:        "RedReader",
		AppPackageName: "org.quantumbadgner.redreader",
		CreatedAt:      strfmt.DateTime(time.Now().Round(0)),
		DateOfUse:      strfmt.Date(time.Now().Round(0)),
		TimeDuration:   int64(time.Minute * 45),
	}
	AppUsage03 = models.AppUsage{
		AppName:        "Xender",
		AppPackageName: "cn.xender",
		CreatedAt:      strfmt.DateTime(time.Now().Round(0)),
		DateOfUse:      strfmt.Date(time.Now().Round(0)),
		TimeDuration:   int64(time.Minute * 45),
	}
)
