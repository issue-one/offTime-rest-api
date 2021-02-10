package mock

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
)

var (
	User01 = models.User{
		Username:    "almaz",
		Email:       "pearl@pax.os",
		Password:    "thejewel",
		PictureURL:  "cam_123124.jpg",
		RoomHistory: []strfmt.UUID{},
	}
	User02 = models.User{
		Username:    "tseahay",
		Email:       "sun@worsh.ip",
		Password:    "kokob",
		PictureURL:  "inyoface.jpg",
		RoomHistory: []strfmt.UUID{},
	}
	User03 = models.User{
		Username:    "emebet",
		Email:       "mistress@nig.ht",
		Password:    "danger",
		PictureURL:  "perplex.jpg",
		RoomHistory: []strfmt.UUID{},
	}
)

func NewMockUserRepositories() repositories.User {
	user01Copy := User01
	user02Copy := User02
	user03Copy := User03
	return repo{
		users: map[string]*models.User{
			User01.Username: &user01Copy,
			User02.Username: &user02Copy,
			User03.Username: &user03Copy,
		},
	}
}

func NewMockUserRepositoriesEmpty() repositories.User {
	return repo{
		users: make(map[string]*models.User),
	}
}

type repo struct {
	users map[string]*models.User
}

func (r repo) CreateUser(ctx context.Context, username string, u *models.CreateUserInput) (*models.User, error) {
	createdAt := time.Now().Round(0)
	newUser := models.User{
		Username:    username,
		Email:       *u.Email,
		Password:    *u.Password,
		CreatedAt:   strfmt.DateTime(createdAt),
		UpdatedAt:   strfmt.DateTime(createdAt),
		PictureURL:  "",
		RoomHistory: []strfmt.UUID{},
	}
	r.users[username] = &newUser
	copy := newUser
	return &copy, nil
}

func (r repo) GetUser(ctx context.Context, username string) (*models.User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	copy := *user
	return &copy, nil
}

func (r repo) GetAllUsers(ctx context.Context, limit int64, offset int64) (items []*models.User, totalCount int, err error) {
	capacity := len(r.users)
	if offset < int64(capacity) {
		capacity = int(offset)
	}
	users := make([]*models.User, 0, limit)
	lastIndex := limit + offset
	var ii int64 = -1
	for _, user := range r.users {
		ii++
		if ii < offset {
			continue
		}
		if ii >= lastIndex {
			break
		}

		copy := *user
		users = append(users, &copy)
	}
	return users, len(r.users), nil
}

func (r repo) UpdateUser(ctx context.Context, username string, u *models.UpdateUserInput) (*models.User, error) {
	oldUser, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	if u.Email != "" {
		oldUser.Email = u.Email
	}
	if u.Password != "" {
		oldUser.Password = u.Password
	}
	copy := *oldUser
	return &copy, nil
}

func (r repo) SetImage(ctx context.Context, username string, imageName string) (*models.User, error) {
	oldUser, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	oldUser.PictureURL = imageName

	copy := *oldUser
	return &copy, nil
}

func (r repo) DeleteUser(ctx context.Context, username string) error {
	delete(r.users, username)
	return nil
}

func (r repo) IsUsernameOccupied(ctx context.Context, username string) (bool, error) {
	_, ok := r.users[username]
	return ok, nil
}

func (r repo) IsEmailOccupied(ctx context.Context, email string, excludedUsername string) (bool, error) {
	for _, user := range r.users {
		if string(user.Email) == email {
			if user.Username == excludedUsername {
				return false, nil
			}
			return true, nil
		}
	}
	return false, nil
}
