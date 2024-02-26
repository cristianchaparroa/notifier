package usecase

import (
	"context"
	"fmt"

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
		fmt.Printf("rate limit error:%v \n", handleErr)
		return handleErr
	}

	sendErr := uc.notifier.Send("Automatic email", n.Recipient, n.Content)
	if sendErr != nil {
		fmt.Printf("send notification error:%v \n", handleErr)
		return sendErr
	}

	return nil
}
