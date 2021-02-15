package tests

/*

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/go-cmp/cmp"
	"github.com/issue-one/offTime-rest-api/gen/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
	. "github.com/issue-one/offTime-rest-api/internal/repositories/mock"
)

func AppUsageRepositoryTestSuite(
	t *testing.T,
	getRepo func() (repositories.AppUsage, error),
	disposeOfRepo func(repo repositories.AppUsage) error) {
	ctx := context.TODO()

	var testEmail strfmt.Email = "email@emai.com"
	var testPassword strfmt.Password = "somepassword"
	testAppUsagename := "someusername"

	tests := []struct {
		name string
		test func(t *testing.T, repo repositories.AppUsage)
	}{
		{
			name: "CreateAppUsage - it succeeds",
			test: func(t *testing.T, repo repositories.AppUsage) {
				user, err := repo.CreateAppUsage(ctx, testAppUsagename, &models.CreateAppUsageInput{
					Email:    &testEmail,
					Password: &testPassword,
				})
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if user.AppUsagename != testAppUsagename {
					t.Errorf("failed: %v != %v", user.AppUsagename, testAppUsagename)
				}
				if string(user.Email) != string(testEmail) {
					t.Errorf("failed: %v != %v", user.AppUsagename, testAppUsagename)
				}
				if string(user.Password) != string(testPassword) {
					t.Errorf("failed: %v != %v", user.AppUsagename, testAppUsagename)
				}
			},
		},
		{
			name: "CreateAppUsage - it assigns CreatedAt",
			test: func(t *testing.T, repo repositories.AppUsage) {
				user, err := repo.CreateAppUsage(ctx, testAppUsagename, &models.CreateAppUsageInput{
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
			name: "CreateAppUsage - it assigns UpdatedAt",
			test: func(t *testing.T, repo repositories.AppUsage) {
				user, err := repo.CreateAppUsage(ctx, testAppUsagename, &models.CreateAppUsageInput{
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
			name: "GetAllAppUsages - it succeeds",
			test: func(t *testing.T, repo repositories.AppUsage) {
				users, _, err := repo.GetAllAppUsages(ctx, 10, 0)
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
			name: "GetAllAppUsages - it returns totalCount",
			test: func(t *testing.T, repo repositories.AppUsage) {
				_, totalCount, err := repo.GetAllAppUsages(ctx, 10, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if totalCount != 3 {
					t.Errorf("totalCount returned %v != 3", totalCount)
				}
			},
		},
		{
			name: "GetAllAppUsages - it limits",
			test: func(t *testing.T, repo repositories.AppUsage) {
				users, _, err := repo.GetAllAppUsages(ctx, 2, 0)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(users) != 2 {
					t.Errorf("limit wasn't respected: %v were returned", len(users))
				}
			},
		},
		{
			name: "GetAllAppUsages - it offsets",
			test: func(t *testing.T, repo repositories.AppUsage) {
				users, _, err := repo.GetAllAppUsages(ctx, 10, 2)
				if err != nil {
					t.Errorf("failed: %v", err)
				}
				if len(users) != 1 {
					t.Errorf("offset wasn't respected: %v were returned", len(users))
				}
			},
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
*/
