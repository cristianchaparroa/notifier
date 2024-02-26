// Package usecase defines the service business logic
package usecase

import (
	"context"
	"notifier/internal/entity"
	"time"
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

// NotificationRateLimitHandler defines the interface for a handler in the chain of responsibility.
type NotificationRateLimitHandler interface {
	Handle(ctx context.Context, n *entity.Notification) error
	SetNext(next NotificationRateLimitHandler)
}

// RateLimitCache defines an interface for storing and retrieving last sent times.
type RateLimitCache interface {
	GetLastSentTime(ctx context.Context, key string) (time.Time, error)
	SetLastSentTime(ctx context.Context, key string, value time.Time) error
}
