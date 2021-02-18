package mock

import (
	"context"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
)

type userRepo struct {
	mutex *sync.RWMutex
	users map[string]*models.User
}

func NewMockUserRepository() repositories.User {
	return userRepo{
		mutex: &sync.RWMutex{},
		users: map[string]*models.User{
			User01.Username: &User01,
			User02.Username: &User02,
			User03.Username: &User03,
		},
	}
}

func NewMockUserRepositoryCopyEntities() repositories.User {
	user01Copy := User01
	user02Copy := User02
	user03Copy := User03
	return userRepo{
		mutex: &sync.RWMutex{},
		users: map[string]*models.User{
			User01.Username: &user01Copy,
			User02.Username: &user02Copy,
			User03.Username: &user03Copy,
		},
	}
}

func NewMockUserRepositoryEmpty() repositories.User {
	return userRepo{
		mutex: &sync.RWMutex{},
		users: make(map[string]*models.User),
	}
}

func (r userRepo) CreateUser(ctx context.Context, username string, u *models.CreateUserInput) (*models.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	createdAt := time.Now().Round(0)
	newUser := models.User{
		Username:    username,
		Email:       *u.Email,
		CreatedAt:   strfmt.DateTime(createdAt),
		UpdatedAt:   strfmt.DateTime(createdAt),
		PictureURL:  "",
		RoomHistory: []strfmt.UUID{},
	}
	r.users[username] = &newUser
	copy := newUser
	return &copy, nil
}

func (r userRepo) GetUser(ctx context.Context, username string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	user, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	// copy := *user
	return user, nil
}

func (r userRepo) GetAllUsers(ctx context.Context, limit int64, offset int64) (items []*models.User, totalCount int, err error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
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

func (r userRepo) UpdateUser(ctx context.Context, username string, u *models.UpdateUserInput) (*models.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	oldUser, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	if u.Email != "" {
		oldUser.Email = u.Email
	}
	copy := *oldUser
	return &copy, nil
}

func (r userRepo) SetImage(ctx context.Context, username string, imageName string) (*models.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	oldUser, ok := r.users[username]
	if !ok {
		return nil, repositories.ErrUserNotFound
	}
	oldUser.PictureURL = imageName

	copy := *oldUser
	return &copy, nil
}

func (r userRepo) DeleteUser(ctx context.Context, username string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.users, username)
	return nil
}

func (r userRepo) IsUsernameOccupied(ctx context.Context, username string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.users[username]
	return ok, nil
}

func (r userRepo) IsEmailOccupied(ctx context.Context, email string, excludedUsername string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
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
