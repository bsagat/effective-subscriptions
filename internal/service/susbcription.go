package service

import (
	"context"
	"log/slog"
	"submanager/internal/domain"
	"submanager/internal/pkg/logger"
	"time"
)

type SubsService struct {
	repo domain.SubsRepo
	log  logger.Logger
}

func NewSubsService(repo domain.SubsRepo, log logger.Logger) *SubsService {
	return &SubsService{
		repo: repo,
		log:  log,
	}
}

// CreateSubscription creates a new subscription in the database.
// It checks if the subscription is unique and sets the expiration date if not provided.
func (s *SubsService) CreateSubscription(ctx context.Context, subs domain.Subscription) error {
	const op = "SubsService.CreateSubscription"
	log := s.log.With(
		slog.String("op", op),
		slog.String("service_name", subs.ServiceName),
		slog.String("user_id", subs.UserID),
		slog.String("start_date", subs.StartDate.String()),
		slog.Int("price", subs.Price),
	)

	if subs.EndDate.IsZero() {
		// Set exp_date as one month from the start of subscription
		subs.EndDate = subs.StartDate.AddDate(0, 1, 0)
		log = log.With("exp_date", subs.EndDate)
	}

	// Check is subscription unique
	unique, err := s.repo.IsUnique(ctx, subs.ServiceName, subs.UserID)
	if err != nil {
		log.Error("Failed to check subscription uniqueness", "error", err)
		return err
	}

	if !unique {
		return domain.ErrSubNotUnique
	}

	// Create a new subscription in the database
	if err := s.repo.Create(ctx, subs); err != nil {
		log.Error("Failed to create new subs", "error", err)
		return err
	}

	log.Info("Subcription has been created")
	return nil
}

// GetSubscription retrieves a subscription by service name and user ID.
func (s *SubsService) GetSubscription(ctx context.Context, serviceName, userID string) (domain.Subscription, error) {
	const op = "SubsService.GetSubscription"
	log := s.log.With(
		slog.String("op", op),
		slog.String("service_name", serviceName),
		slog.String("user_ID", userID),
	)

	subs, err := s.repo.Get(ctx, serviceName, userID)
	if err != nil {
		log.Error("Failed to get subscription", "error", err)
		return domain.Subscription{}, err
	}

	log.Info("Subscription has been retrieved")
	return subs, nil
}

// GetSubscriptionList retrieves all subscriptions for a given user ID.
func (s *SubsService) GetSubscriptionList(ctx context.Context, userID string) ([]domain.Subscription, error) {
	const op = "SubsService.GetSubscriptionList"
	log := s.log.With(
		slog.String("op", op),
		slog.String("user_ID", userID),
	)

	subs, err := s.repo.List(ctx, userID)
	if err != nil {
		log.Error("Failed to get subscription list", "error", err)
		return nil, err
	}

	if len(subs) == 0 {
		log.Error("Subscription list is empty")
		return nil, domain.ErrSubsNotFound
	}

	log.Info("Subscription list has been retrieved")
	return subs, nil
}

// UpdateSubscription updates an existing subscription in the database.
func (s *SubsService) UpdateSubscription(ctx context.Context, subs domain.Subscription) error {
	const op = "SubsService.UpdateSubs"
	log := s.log.With(
		slog.String("op", op),
		slog.String("service_name", subs.ServiceName),
		slog.String("user_id", subs.UserID),
		slog.String("start_date", subs.StartDate.String()),
		slog.Int("price", subs.Price),
	)

	if subs.EndDate.IsZero() {
		// Set exp_date as one month from the start of subscription
		subs.EndDate = subs.StartDate.AddDate(0, 1, 0)
		log = log.With("exp_date", subs.EndDate)
	}

	if err := s.repo.Update(ctx, subs); err != nil {
		log.Error("Failed to update subscription", "error", err)
		return err
	}

	log.Info("Subscription has been updated")
	return nil
}

// DeleteSubscription deletes a subscription by service name and user ID.
func (s *SubsService) DeleteSubscription(ctx context.Context, serviceName, userID string) error {
	const op = "SubsService.DeleteSubscription"
	log := s.log.With(
		slog.String("op", op),
		slog.String("service_name", serviceName),
		slog.String("user_ID", userID),
	)

	if err := s.repo.Delete(ctx, serviceName, userID); err != nil {
		log.Error("Failed to delete subscription", "error", err)
		return err
	}

	log.Info("Subscription has been deleted")
	return nil
}

// DeleteSubscriptionList deletes all subscriptions for a given user ID.
func (s *SubsService) DeleteSubscriptionList(ctx context.Context, userID string) error {
	const op = "SubsService.DeleteSubscription"
	log := s.log.With(
		slog.String("op", op),
		slog.String("user_ID", userID),
	)

	if err := s.repo.DeleteList(ctx, userID); err != nil {
		log.Error("Failed to delete subscription list", "error", err)
		return err
	}

	log.Info("Subscription list has been deleted")
	return nil
}

// GetSummaryByFilter retrieves a summary of subscriptions based on the provided filter criteria.
// It returns the total price, count of subscriptions, and the list of subscriptions.
func (s *SubsService) GetSummaryByFilter(ctx context.Context, start, end time.Time, serviceName, userID string, pageNum, pageSize int) (domain.Summary, error) {
	const op = "SubsService.GetSummaryByFilter"
	log := s.log.With(
		slog.String("op", op),
		slog.String("service_name", serviceName),
		slog.String("user_id", userID),
		slog.String("filter_start_date", start.String()),
		slog.String("filter_end_date", end.String()),
		slog.Int("page_number", pageNum),
		slog.Int("page_size", pageSize),
	)

	subs, err := s.repo.SubsListByFilter(ctx, start, end, serviceName, userID, pageNum, pageSize)
	if err != nil {
		log.Error("Failed to get subs list by filter", "error", err)
		return domain.Summary{}, err
	}

	if len(subs) == 0 {
		log.Error("Subscription list is empty")
		return domain.Summary{}, domain.ErrSubsNotFound
	}

	summary := domain.Summary{
		TotalPrice:    getSummary(subs),
		SubsCount:     len(subs),
		PageNumber:    pageNum,
		PageSize:      pageSize,
		Subscriptions: subs,
	}
	log.Info("Subscription list summary has been calculated successfully", "subs_count", summary.SubsCount, "total_price", summary.TotalPrice, "page_number", summary.PageNumber, "page_size", summary.PageSize)
	return summary, nil
}

func getSummary(subsList []domain.Subscription) int {
	var total int
	for _, sub := range subsList {
		total += sub.Price
	}
	return total
}
