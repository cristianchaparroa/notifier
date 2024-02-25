package usecase

import (
	"context"
	"notifier/internal/entity"
)

type notificationUseCase struct {
	notifier EmailManager
}

func NewNotificationUseCase(notifier EmailManager) NotificationUseCase {
	return &notificationUseCase{
		notifier: notifier,
	}
}

func (uc *notificationUseCase) Notify(ctx context.Context, n *entity.Notification) error {
	panic("implement me")
}
