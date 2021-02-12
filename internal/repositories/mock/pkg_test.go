package mock_test

import (
	"testing"

	"github.com/issue-one/offTime-rest-api/internal/repositories"
	"github.com/issue-one/offTime-rest-api/internal/repositories/mock"
	"github.com/issue-one/offTime-rest-api/internal/repositories/tests"
)

func TestUserRepository(t *testing.T) {
	tests.UserRepositoryTestSuite(t,
		func() (repositories.User, error) {
			return mock.NewMockUserRepositoriesCopyEntities(), nil
		},
		func(repo repositories.User) error {
			return nil
		},
	)
}

func TestRoomRepository(t *testing.T) {
	tests.RoomRepositoryTestSuite(t,
		func() (repositories.Room, error) {
			return mock.NewMockRoomRepositoriesCopyEntities(), nil
		},
		func(repo repositories.Room) error {
			return nil
		},
	)
}
