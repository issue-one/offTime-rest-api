package mock

import (
	"context"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/issue-one/offTime-rest-api/gen/models"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
)

type appUsageRepo struct {
	mutex  *sync.RWMutex
	usages map[string]map[string]*models.AppUsage
}

func AppUsageId(u *models.AppUsage) string {
	return u.AppPackageName + u.DateOfUse.String()
}

func NewMockAppUsageRepository() repositories.AppUsage {
	return appUsageRepo{
		mutex: &sync.RWMutex{},
		usages: map[string]map[string]*models.AppUsage{
			User01.Username: map[string]*models.AppUsage{
				AppUsageId(&AppUsage01): &AppUsage01,
			},
			User02.Username: map[string]*models.AppUsage{
				AppUsageId(&AppUsage02): &AppUsage02,
			},
			User03.Username: map[string]*models.AppUsage{
				AppUsageId(&AppUsage03): &AppUsage03,
			},
		},
	}
}

func NewMockAppUsageRepositoryCopyEntities() repositories.AppUsage {
	usage01Copy := AppUsage01
	usage02Copy := AppUsage02
	usage03Copy := AppUsage03
	return appUsageRepo{
		mutex: &sync.RWMutex{},
		usages: map[string]map[string]*models.AppUsage{
			User01.Username: map[string]*models.AppUsage{
				AppUsageId(&usage01Copy): &usage01Copy,
			},
			User02.Username: map[string]*models.AppUsage{
				AppUsageId(&usage02Copy): &usage02Copy,
			},
			User03.Username: map[string]*models.AppUsage{
				AppUsageId(&usage03Copy): &usage03Copy,
			},
		},
	}
}

func NewMockAppUsageRepositoryEmpty() repositories.AppUsage {
	return appUsageRepo{
		mutex:  &sync.RWMutex{},
		usages: make(map[string]map[string]*models.AppUsage),
	}
}

func (r appUsageRepo) CreateAppUsage(ctx context.Context, username string, u *models.CreateAppUsageInput) (*models.AppUsage, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	createdAt := time.Now().Round(0)
	newAppUsage := models.AppUsage{
		AppName:        *u.AppName,
		AppPackageName: *u.AppPackageName,
		DateOfUse:      *u.DateOfUse,
		TimeDuration:   *u.TimeDuration,
		CreatedAt:      strfmt.DateTime(createdAt),
	}
	usages, ok := r.usages[username]
	if !ok {
		usages = make(map[string]*models.AppUsage)
		r.usages[username] = usages
	}
	usages[AppUsageId(&newAppUsage)] = &newAppUsage
	copy := newAppUsage
	return &copy, nil
}

func (r appUsageRepo) GetAllAppUsages(ctx context.Context, username string, limit, offset int64) (items []*models.AppUsage, totalCount int, err error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	capacity := len(r.usages)
	if offset < int64(capacity) {
		capacity = int(offset)
	}
	appUsages := make([]*models.AppUsage, 0, limit)

	userUsages, ok := r.usages[username]
	if !ok {
		return appUsages, 0, nil
	}
	lastIndex := limit + offset
	var ii int64 = -1
	for _, usages := range userUsages {
		ii++
		if ii < offset {
			continue
		}
		if ii >= lastIndex {
			break
		}

		copy := *usages
		appUsages = append(appUsages, &copy)
	}
	return appUsages, len(userUsages), nil
}

func (r appUsageRepo) DeleteAppUsages(ctx context.Context, username string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.usages, username)
	return nil
}
