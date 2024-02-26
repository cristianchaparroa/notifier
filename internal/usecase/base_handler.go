package usecase

import (
	"context"
	"fmt"
	"notifier/internal/entity"
	"time"
)

type RateLimitRule struct {
	MaxPerUnit int
	Unit       time.Duration
}

// BaseHandler handles rate limits for the "any" notification type.
type BaseHandler struct {
	MaxPerUnit int                          // Maximum number of notifications allowed per unit (e.g., 2)
	Unit       time.Duration                // Unit of time for the rate limit (e.g., time.Minute)
	Next       NotificationRateLimitHandler // Next handler in the chain
	cache      RateLimitCache
}

func (h *BaseHandler) Validate(ctx context.Context, notification *entity.Notification) error {
	lastSentTime, getErr := h.cache.GetLastSentTime(ctx, notification.Recipient)
	if getErr != nil {
		return getErr
	}

	if time.Since(lastSentTime) >= h.Unit {
		h.cache.SetLastSentTime(ctx, notification.Recipient, time.Now())
		return nil
	}

	return fmt.Errorf("notification rate limit exceeded for type %q: recipient %q", notification.Type, notification.Recipient)
}

func BuildRateLimitChain(cache RateLimitCache) NotificationRateLimitHandler {
	// TODO: it could be stored on db to be configurable
	rules := map[string]RateLimitRule{
		entity.StatusNotificationType:    {MaxPerUnit: 2, Unit: time.Minute},
		entity.NewsNotificationType:      {MaxPerUnit: 1, Unit: time.Hour * 24},
		entity.MarketingNotificationType: {MaxPerUnit: 3, Unit: time.Hour},
	}

	sh := NewStatusHandler(cache, rules[entity.StatusNotificationType])
	nh := NewNewsHandler(cache, rules[entity.NewsNotificationType])
	mh := NewMarketingHandler(cache, rules[entity.MarketingNotificationType])
	uh := &UnknownHandler{}

	sh.Next = nh
	nh.Next = mh
	mh.Next = uh // the UnknownHandler should be the last one

	return sh
}
