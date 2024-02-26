package usecase

import (
	"fmt"
	"notifier/internal/entity"
	"sync"
	"time"
)

type RateLimitRule struct {
	MaxPerUnit int
	Unit       time.Duration
}

// BaseHandler handles rate limits for the "any" notification type.
type BaseHandler struct {
	MaxPerUnit    int                          // Maximum number of notifications allowed per unit (e.g., 2)
	Unit          time.Duration                // Unit of time for the rate limit (e.g., time.Minute)
	lastSentTimes map[string]time.Time         // Maps recipient emails to the last time sent
	mutex         *sync.Mutex                  // Protects lastSentTimes from concurrent access
	Next          NotificationRateLimitHandler // Next handler in the chain
}

func (h *BaseHandler) Validate(notification *entity.Notification) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	lastSentTime, ok := h.lastSentTimes[notification.Recipient]
	if !ok || time.Since(lastSentTime) >= h.Unit {
		h.lastSentTimes[notification.Recipient] = time.Now()
		return nil
	}

	return fmt.Errorf("notification rate limit exceeded for type %q: recipient %q", notification.Type, notification.Recipient)
}

func BuildRateLimitChain() NotificationRateLimitHandler {
	// TODO: it could be stored on db to be configurable
	rules := map[string]RateLimitRule{
		entity.StatusNotificationType:    {MaxPerUnit: 2, Unit: time.Minute},
		entity.NewsNotificationType:      {MaxPerUnit: 1, Unit: time.Hour * 24},
		entity.MarketingNotificationType: {MaxPerUnit: 3, Unit: time.Hour},
	}

	// TODO: abstract the storage to be scalable
	// 	i.e: Redis, Memcache...(ephemeral storage with high response)
	lastSentTimes := make(map[string]time.Time)
	m := &sync.Mutex{}

	sh := NewStatusHandler(lastSentTimes, rules[entity.StatusNotificationType], m)
	nh := NewNewsHandler(lastSentTimes, rules[entity.NewsNotificationType], m)
	mh := NewMarketingHandler(lastSentTimes, rules[entity.MarketingNotificationType], m)
	uh := &UnknownHandler{}

	sh.Next = nh
	nh.Next = mh
	mh.Next = uh // the UnknownHandler should be the last one

	return sh
}
