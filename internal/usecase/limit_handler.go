package usecase

import (
	"context"
	"fmt"
	"notifier/internal/entity"
)

// StatusHandler handles rate limits for the "status" notification type.
type StatusHandler struct {
	*BaseHandler
}

func NewStatusHandler(cache RateLimitCache, r RateLimitRule) *StatusHandler {
	return &StatusHandler{&BaseHandler{
		MaxPerUnit: r.MaxPerUnit,
		Unit:       r.Unit,
		cache:      cache,
	}}
}

func (h *StatusHandler) Handle(ctx context.Context, n *entity.Notification) error {
	if entity.StatusNotificationType != n.Type {
		return h.Next.Handle(ctx, n) // Pass to Next handler
	}
	return h.Validate(ctx, n)
}

func (h *StatusHandler) SetNext(next NotificationRateLimitHandler) {
	h.Next = next
}

// MarketingHandler handles rate limits for the "marketing" notification type.
type MarketingHandler struct {
	*BaseHandler
}

func NewMarketingHandler(cache RateLimitCache, r RateLimitRule) *MarketingHandler {
	return &MarketingHandler{&BaseHandler{
		MaxPerUnit: r.MaxPerUnit,
		Unit:       r.Unit,
		cache:      cache,
	}}
}

func (h *MarketingHandler) Handle(ctx context.Context, n *entity.Notification) error {
	if entity.MarketingNotificationType != n.Type {
		return h.Next.Handle(ctx, n) // Pass to Next handler
	}
	return h.Validate(ctx, n)
}

func (h *MarketingHandler) SetNext(next NotificationRateLimitHandler) {
	h.Next = next
}

// NewsHandler handles rate limits for the "news" notification type.
type NewsHandler struct {
	*BaseHandler
}

func NewNewsHandler(cache RateLimitCache, r RateLimitRule) *NewsHandler {
	return &NewsHandler{&BaseHandler{
		MaxPerUnit: r.MaxPerUnit,
		Unit:       r.Unit,
		cache:      cache,
	}}
}

func (h *NewsHandler) Handle(ctx context.Context, n *entity.Notification) error {
	if entity.NewsNotificationType != n.Type {
		return h.Next.Handle(ctx, n) // Pass to Next handler
	}
	return h.Validate(ctx, n)
}

func (h *NewsHandler) SetNext(next NotificationRateLimitHandler) {
	h.Next = next
}

type UnknownHandler struct {
	next NotificationRateLimitHandler
}

func (h *UnknownHandler) Handle(ctx context.Context, n *entity.Notification) error {
	return fmt.Errorf("notification rate limit unhandled for type %q: recipient %q", n.Type, n.Recipient)
}

func (h *UnknownHandler) SetNext(next NotificationRateLimitHandler) {
	h.next = next
}
