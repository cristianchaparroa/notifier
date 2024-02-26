package usecase

import (
	"fmt"
	"notifier/internal/entity"
	"sync"
	"time"
)

// StatusHandler handles rate limits for the "status" notification type.
type StatusHandler struct {
	*BaseHandler
}

func NewStatusHandler(lastSentTimes map[string]time.Time, r RateLimitRule, m sync.Mutex) *StatusHandler {
	return &StatusHandler{&BaseHandler{
		MaxPerUnit:    r.MaxPerUnit,
		Unit:          r.Unit,
		lastSentTimes: lastSentTimes,
		mutex:         m,
	}}
}

func (h *StatusHandler) Handle(n *entity.Notification) error {
	if entity.StatusNotificationType != n.Type {
		return h.next.Handle(n) // Pass to next handler
	}
	return h.Validate(n)
}

func (h *StatusHandler) SetNext(next NotificationRateLimitHandler) {
	h.next = next
}

// MarketingHandler handles rate limits for the "marketing" notification type.
type MarketingHandler struct {
	*BaseHandler
}

func NewMarketingHandler(lastSentTimes map[string]time.Time, r RateLimitRule, m sync.Mutex) *MarketingHandler {
	return &MarketingHandler{&BaseHandler{
		MaxPerUnit:    r.MaxPerUnit,
		Unit:          r.Unit,
		lastSentTimes: lastSentTimes,
		mutex:         m,
	}}
}

func (h *MarketingHandler) Handle(n *entity.Notification) error {
	if entity.MarketingNotificationType != n.Type {
		return h.next.Handle(n) // Pass to next handler
	}
	return h.Validate(n)
}

func (h *MarketingHandler) SetNext(next NotificationRateLimitHandler) {
	h.next = next
}

// NewsHandler handles rate limits for the "news" notification type.
type NewsHandler struct {
	*BaseHandler
}

func NewNewsHandler(lastSentTimes map[string]time.Time, r RateLimitRule, m sync.Mutex) *NewsHandler {
	return &NewsHandler{&BaseHandler{
		MaxPerUnit:    r.MaxPerUnit,
		Unit:          r.Unit,
		lastSentTimes: lastSentTimes,
		mutex:         m,
	}}
}

func (h *NewsHandler) Handle(n *entity.Notification) error {
	if entity.NewsNotificationType != n.Type {
		return h.next.Handle(n) // Pass to next handler
	}
	return h.Validate(n)
}

func (h *NewsHandler) SetNext(next NotificationRateLimitHandler) {
	h.next = next
}

type UnknownHandler struct {
	next NotificationRateLimitHandler
}

func (h *UnknownHandler) Handle(n *entity.Notification) error {
	return fmt.Errorf("notification rate limit unhandled for type %q: recipient %q", n.Type, n.Recipient)
}

func (h *UnknownHandler) SetNext(next NotificationRateLimitHandler) {
	h.next = next
}
