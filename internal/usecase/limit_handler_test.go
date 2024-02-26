package usecase

import (
	"fmt"
	"notifier/internal/entity"
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

func (s *limitHandlerSuite) setupTest() {}

func TestLimitHandlerSuite(t *testing.T) {
	suite.Run(t, new(limitHandlerSuite))
}

func (s *limitHandlerSuite) TestChainOfResponsibility() {
	handler := BuildRateLimitChain()

	// Sending notifications (some will be rejected due to rate limits)
	notifications := []notificationTest{
		{entity.NewNotification("", entity.StatusNotificationType, "user1@example.com"), true},
		{entity.NewNotification("", entity.StatusNotificationType, "user1@example.com"), false}, // Second one will be rejected
		{entity.NewNotification("", entity.NewsNotificationType, "user2@example.com"), true},
		{entity.NewNotification("", entity.MarketingNotificationType, "user3@example.com"), true},
		{entity.NewNotification("", entity.MarketingNotificationType, "user3@example.com"), false}, // Second one will be rejected
		{entity.NewNotification("", "unknown", "user4@example.com"), false},                        // Not Allowed (no handler for "unknown")
	}

	for i, tc := range notifications {
		s.T().Run(fmt.Sprintf("notification-%d", i), func(t *testing.T) {
			err := handler.Handle(tc.n)
			if !tc.expectedAllowed {
				s.NotNil(err)
			}
		})
	}
}
