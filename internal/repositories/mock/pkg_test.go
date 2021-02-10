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
			return mock.NewMockUserRepositories(), nil
		},
		func(repo repositories.User) error {
			return nil
		},
	)
}
