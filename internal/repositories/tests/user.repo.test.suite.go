package tests

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/go-cmp/cmp"
	"github.com/issue-one/offTime-rest-api/gen/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
	. "github.com/issue-one/offTime-rest-api/internal/repositories/mock"
)

func UserRepositoryTestSuite(
	t *testing.T,
	getRepo func() (repositories.User, error),
	disposeOfRepo func(repo repositories.User) error) {
	ctx := context.TODO()

	var testEmail strfmt.Email = "email@emai.com"
	var testPassword strfmt.Password = "somepassword"
	testUsername := "someusername"

	tests := []struct {
		name string
		test func(t *testing.T, repo repositories.User)
	}{
		{
			name: "GetUser - it succeeds",
			test: func(t *testing.T, repo repositories.User) {
				user, err := repo.GetUser(ctx, User01.Username)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if user.Username != User01.Username {
					t.Errorf("failed: %v != %v", user.Username, User01.Username)
				}
				if user.Email != strfmt.Email(User01.Email) {
					t.Errorf("failed: %v != %v", user.Email, User01.Email)
				}
				if user.PictureURL != User01.PictureURL {
					t.Errorf("failed: %v != %v", user.PictureURL, User01.PictureURL)
				}
				if !cmp.Equal(user.RoomHistory, User01.RoomHistory) {
					t.Errorf("failed: %v != %v", user.RoomHistory, User01.RoomHistory)
				}
			},
		},
		{
			name: "GetUser - it throws ErrUserNotFound if User not found",
			test: func(t *testing.T, repo repositories.User) {
				_, err := repo.GetUser(ctx, "fakeUsername")
				if err == nil || err != repositories.ErrUserNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
		{
			name: "CreateUser - it succeeds",
			test: func(t *testing.T, repo repositories.User) {
				user, err := repo.CreateUser(ctx, testUsername, &models.CreateUserInput{
					Email:    &testEmail,
					Password: &testPassword,
				})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if user.Username != testUsername {
					t.Errorf("failed: %v != %v", user.Username, testUsername)
				}
				if string(user.Email) != string(testEmail) {
					t.Errorf("failed: %v != %v", user.Username, testUsername)
				}
			},
		},
		{
			name: "CreateUser - it assigns CreatedAt",
			test: func(t *testing.T, repo repositories.User) {
				user, err := repo.CreateUser(ctx, testUsername, &models.CreateUserInput{
					Email:    &testEmail,
					Password: &testPassword,
				})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if user.CreatedAt.Equal(strfmt.DateTime{}) {
					t.Errorf("failed: cratedAt is nil\n%v", user)
				}
			},
		},
		{
			name: "CreateUser - it assigns UpdatedAt",
			test: func(t *testing.T, repo repositories.User) {
				user, err := repo.CreateUser(ctx, testUsername, &models.CreateUserInput{
					Email:    &testEmail,
					Password: &testPassword,
				})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if user.UpdatedAt.Equal(strfmt.DateTime{}) {
					t.Errorf("failed: cratedAt is nil\n%v", user)
				}
			},
		},
		{
			name: "UpdateUser - it succeeds",
			test: func(t *testing.T, repo repositories.User) {
				user, err := repo.UpdateUser(ctx, User01.Username, &models.UpdateUserInput{
					Email:    testEmail,
					Password: testPassword,
				})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if string(user.Email) != string(testEmail) {
					t.Errorf("failed: %v != %v", user.Username, testUsername)
				}
			},
		},
		{
			name: "UpdateUser - it throws ErrUserNotFound if User not found",
			test: func(t *testing.T, repo repositories.User) {
				_, err := repo.UpdateUser(ctx, "fakeUsername", &models.UpdateUserInput{
					Email:    testEmail,
					Password: testPassword,
				})
				if err == nil || err != repositories.ErrUserNotFound {
					t.Errorf("failed: err == %v", err)
				}
			},
		},
		{
			name: "GetAllUsers - it succeeds",
			test: func(t *testing.T, repo repositories.User) {
				users, _, err := repo.GetAllUsers(ctx, 10, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if users == nil {
					t.Error("nil array returned")
					return
				}
				if len(users) != 3 {
					t.Errorf("only %v were returned", len(users))
				}
			},
		},
		{
			name: "GetAllUsers - it returns totalCount",
			test: func(t *testing.T, repo repositories.User) {
				_, totalCount, err := repo.GetAllUsers(ctx, 10, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if totalCount != 3 {
					t.Errorf("totalCount returned %v != 3", totalCount)
				}
			},
		},
		{
			name: "GetAllUsers - it limits",
			test: func(t *testing.T, repo repositories.User) {
				users, _, err := repo.GetAllUsers(ctx, 2, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(users) != 2 {
					t.Errorf("limit wasn't respected: %v were returned", len(users))
				}
			},
		},
		{
			name: "GetAllUsers - it offsets",
			test: func(t *testing.T, repo repositories.User) {
				users, _, err := repo.GetAllUsers(ctx, 10, 2)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(users) != 1 {
					t.Errorf("offset wasn't respected: %v were returned", len(users))
				}
			},
		},
		{
			name: "IsEmailOccupied - it detects occupied",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsEmailOccupied(ctx, User01.Email.String(), "")
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if !isOccupied {
					t.Errorf("occupied email reported as unoccupied")
				}
			},
		},
		{
			name: "IsEmailOccupied - it detects unoccupied",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsEmailOccupied(ctx, "some@email.com", "")
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if isOccupied {
					t.Errorf("occupied email reported as unoccupied")
				}
			},
		},
		{
			name: "IsEmailOccupied - it accepts exclusion",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsEmailOccupied(ctx, User01.Email.String(), User01.Username)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if isOccupied {
					t.Errorf("non-occupied email reported as occupied")
				}
			},
		},
		{
			name: "IsUsernameOccupied - it detects occupied",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsUsernameOccupied(ctx, User01.Username)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if !isOccupied {
					t.Error("occupied email reported as unoccupied")
				}
			},
		},
		{
			name: "IsUsernameOccupied - it detects unoccupied",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsUsernameOccupied(ctx, User01.Email.String())
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if isOccupied {
					t.Errorf("occupied email reported as unoccupied")
				}
			},
		},
		{
			name: "IsUsernameOccupied - it accepts exclusion",
			test: func(t *testing.T, repo repositories.User) {
				isOccupied, err := repo.IsUsernameOccupied(ctx, User01.Email.String())
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if isOccupied {
					t.Errorf("non-occupied email reported as occupied")
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
