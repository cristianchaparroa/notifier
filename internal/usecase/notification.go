package usecase

import (
	"context"

	"notifier/internal/entity"
)

type notificationUseCase struct {
	notifier         EmailManager
	rateLimitHandler NotificationRateLimitHandler
}

func NewNotificationUseCase(notifier EmailManager, h NotificationRateLimitHandler) NotificationUseCase {
	return &notificationUseCase{
		notifier:         notifier,
		rateLimitHandler: h,
	}
}

func (uc *notificationUseCase) Notify(ctx context.Context, n *entity.Notification) error {
	handleErr := uc.rateLimitHandler.Handle(n)
	if handleErr != nil {
		return handleErr
	}

	sendErr := uc.notifier.Send("Automatic email", n.Recipient, n.Content)
	if sendErr != nil {
		return sendErr
	}

	return nil
}
