package mock

import (
	"context"
	"submanager/internal/core/domain"
	"time"
)

type MockSubsRepo struct {
}

func NewMockSubsRepo() *MockSubsRepo {
	return &MockSubsRepo{}
}

func (repo *MockSubsRepo) Create(ctx context.Context, subs domain.Subscription) error {
	return nil
}
func (repo *MockSubsRepo) Delete(ctx context.Context, serviceName string, userID string) error {
	if serviceName == "notexist" {
		return domain.ErrSubsNotFound
	}
	return nil
}
func (repo *MockSubsRepo) DeleteList(ctx context.Context, userID string) error {
	if userID == "notexist" {
		return domain.ErrSubsNotFound
	}
	return nil
}
func (repo *MockSubsRepo) Get(ctx context.Context, serviceName string, userID string) (domain.Subscription, error) {
	if serviceName == "notexist" {
		return domain.Subscription{}, domain.ErrSubsNotFound
	}
	return domain.Subscription{}, nil
}
func (repo *MockSubsRepo) IsUnique(ctx context.Context, serviceName string, userID string) (bool, error) {
	if serviceName == "notunique" {
		return false, nil
	}
	return true, nil
}
func (repo *MockSubsRepo) List(ctx context.Context, userID string) ([]domain.Subscription, error) {
	if userID == "notexist" {
		return []domain.Subscription{}, nil
	}
	return []domain.Subscription{
		{},
	}, nil
}
func (repo *MockSubsRepo) SubsListByFilter(ctx context.Context, start time.Time, end time.Time, serviceName, userID string, pageNum, pageSize int) ([]domain.Subscription, error) {
	if serviceName == "notexist" {
		return []domain.Subscription{}, nil
	}
	return []domain.Subscription{
		{},
	}, nil
}
func (repo *MockSubsRepo) Update(ctx context.Context, subs domain.Subscription) error {
	if subs.ServiceName == "notexist" {
		return domain.ErrSubsNotFound
	}
	return nil
}
