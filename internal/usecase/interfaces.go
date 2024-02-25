package usecase

import (
	"context"
	"notifier/internal/entity"
)

// NotificationUseCase defines the logic that manages and sends out email notifications of various types.
type NotificationUseCase interface {
	Notify(ctx context.Context, n *entity.Notification) error
}

// EmailManager is in charge to send emails
type EmailManager interface {
	// Send will send emails according  to the following params
	// addressee is the email receiver
	// body is the message content
	Send(subject, addressee, body string) error
}
