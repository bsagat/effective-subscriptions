package service

import (
	"context"
	"log/slog"
	"submanager/internal/adapters/repo"
	"submanager/internal/domain"
)

type SubsService struct {
	repo *repo.SubsRepo
	log  *slog.Logger
}

func NewSubsService(repo *repo.SubsRepo, log *slog.Logger) *SubsService {
	return &SubsService{
		repo: repo,
		log:  log,
	}
}

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
		log.Error("Failed to create new subs in DB", "error", err)
		return err
	}

	log.Info("Subcription has been created")
	return nil
}

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
