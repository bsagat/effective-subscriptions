package domain

import (
	"context"
	"time"
)

// ---------------- Subs Repository ----------------

type SubsRepo interface {
	SubsCreator
	SubsUpdater
	SubsDeleter
	SubsGetter
	SubsChecker
}

type SubsCreator interface {
	Create(ctx context.Context, subs Subscription) error
}

type SubsUpdater interface {
	Update(ctx context.Context, subs Subscription) error
}

type SubsDeleter interface {
	Delete(ctx context.Context, serviceName string, userID string) error
	DeleteList(ctx context.Context, userID string) error
}

type SubsGetter interface {
	Get(ctx context.Context, serviceName string, userID string) (Subscription, error)
	List(ctx context.Context, userID string) ([]Subscription, error)
	SubsListByFilter(ctx context.Context, start time.Time, end time.Time, serviceName string, userID string) ([]Subscription, error)
}

type SubsChecker interface {
	IsUnique(ctx context.Context, serviceName string, userID string) (bool, error)
}

// ---------------- Subs Service ----------------

type SubsService interface {
	CreateSubscription(ctx context.Context, subs Subscription) error
	DeleteSubscription(ctx context.Context, serviceName string, userID string) error
	DeleteSubscriptionList(ctx context.Context, userID string) error
	GetSubscription(ctx context.Context, serviceName string, userID string) (Subscription, error)
	GetSubscriptionList(ctx context.Context, userID string) ([]Subscription, error)
	UpdateSubscription(ctx context.Context, subs Subscription) error
	SummaryService
}

type SummaryService interface {
	GetSummaryByFilter(ctx context.Context, start time.Time, end time.Time, serviceName string, userID string) (Summary, error)
}
