package usecase

import (
	"context"
	"notifier/internal/entity"
	"notifier/internal/usecase/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type notificationTest struct {
	n               *entity.Notification
	expectedAllowed bool
}

type limitHandlerSuite struct {
	suite.Suite
}

func TestLimitHandlerSuite(t *testing.T) {
	suite.Run(t, new(limitHandlerSuite))
}

func (s *limitHandlerSuite) TestChainOfResponsibility() {
	memoryCache := repo.NewInMemoryCache()
	handler := BuildRateLimitChain(memoryCache)

	// Sending notifications (some will be rejected due to rate limits)
	notifications := []notificationTest{
		{entity.NewNotification("", entity.StatusNotificationType, "user1@example.com"), true},
		{entity.NewNotification("", entity.StatusNotificationType, "user1@example.com"), false}, // Second one will be rejected
		{entity.NewNotification("", entity.NewsNotificationType, "user2@example.com"), true},
		{entity.NewNotification("", entity.MarketingNotificationType, "user3@example.com"), true},
		{entity.NewNotification("", entity.MarketingNotificationType, "user3@example.com"), false}, // Second one will be rejected
		{entity.NewNotification("", "unknown", "user4@example.com"), false},                        // Not Allowed (no handler for "unknown")
	}

	ctx := context.Background()
	for _, tc := range notifications {
		err := handler.Handle(ctx, tc.n)
		if !tc.expectedAllowed {
			s.NotNil(err)
		}
	}
}
