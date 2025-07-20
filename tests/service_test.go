package tests

import (
	"context"
	"log/slog"
	"os"
	"submanager/internal/domain"
	"submanager/internal/pkg/logger"
	"submanager/internal/service"
	mock "submanager/tests/mocks"
	"testing"
	"time"
)

var (
	serv *service.SubsService
)

func TestMain(m *testing.M) {
	slog.Info("Starting the tests...")

	// Initialize the mock repository and service
	repo := mock.NewMockSubsRepo()
	log := logger.New(logger.Debug)
	serv = service.NewSubsService(repo, log)

	defer os.Exit(m.Run())
	slog.Info("Test has been finished...")
}

func TestCreateSubscription(t *testing.T) {
	ctx := context.Background()

	// Define a sample subscription
	sampleSubscription := domain.Subscription{
		ServiceName: "TestService",
		UserID:      "user123",
		StartDate:   time.Now(),
		Price:       100,
	}

	// Default test case
	err := serv.CreateSubscription(ctx, sampleSubscription)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not unique
	sampleSubscription.ServiceName = "notunique"
	if err := serv.CreateSubscription(ctx, sampleSubscription); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubNotUnique, err)
	}
}

func TestGetSubscription(t *testing.T) {
	ctx := context.Background()

	serviceName, userID := "TestService", "user123"

	// Default test case
	_, err := serv.GetSubscription(ctx, serviceName, userID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not found
	serviceName = "notexist"
	_, err = serv.GetSubscription(ctx, serviceName, userID)
	if err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}

func TestGetSubscriptionList(t *testing.T) {
	ctx := context.Background()
	userId := "user123"

	// Default test case
	if _, err := serv.GetSubscriptionList(ctx, userId); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription list not found
	userId = "notexist"
	if _, err := serv.GetSubscriptionList(ctx, userId); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}

func TestUpdateSubscription(t *testing.T) {
	ctx := context.Background()

	// Define a sample subscription
	sampleSubscription := domain.Subscription{
		ServiceName: "TestService",
		UserID:      "user123",
		StartDate:   time.Now(),
		Price:       100,
	}

	// Default test case
	if err := serv.UpdateSubscription(ctx, sampleSubscription); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not found
	sampleSubscription.ServiceName = "notexist"
	if err := serv.UpdateSubscription(ctx, sampleSubscription); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}

func TestDeleteSubscription(t *testing.T) {
	ctx := context.Background()

	// Default test case
	serviceName, userID := "TestService", "user123"
	if err := serv.DeleteSubscription(ctx, serviceName, userID); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not found
	serviceName = "notexist"
	if err := serv.DeleteSubscription(ctx, serviceName, userID); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}

func TestDeleteSubscriptionList(t *testing.T) {
	ctx := context.Background()

	// Default test case
	userID := "user123"
	if err := serv.DeleteSubscriptionList(ctx, userID); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not found
	userID = "notexist"
	if err := serv.DeleteSubscriptionList(ctx, userID); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}

func TestGetSummaryByFilter(t *testing.T) {
	ctx := context.Background()

	// Default test case
	start, end, serviceName, userID, pageNum, pageSize := time.Now(), time.Now(), "TestService", "user123", 1, 10
	if _, err := serv.GetSummaryByFilter(ctx, start, end, serviceName, userID, pageNum, pageSize); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if subscription not found
	serviceName = "notexist"
	if _, err := serv.GetSummaryByFilter(ctx, start, end, serviceName, userID, pageNum, pageSize); err == nil {
		t.Errorf("Expected error %v, got %v", domain.ErrSubsNotFound, err)
	}
}
