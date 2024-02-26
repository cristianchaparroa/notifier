package v1

import (
	"context"
	"net/http"
	"notifier/internal/entity"
	"notifier/internal/usecase"

	"github.com/labstack/echo/v4"
)

type NotificationController struct {
	uc usecase.NotificationUseCase
}

func NewNotificationController(uc usecase.NotificationUseCase) *NotificationController {
	return &NotificationController{
		uc: uc,
	}
}

func (r *NotificationController) SendNotifications(c echo.Context) error {
	ctx := context.Background()
	n, decodeErr := entity.NewNotificationFromRequest(c.Request().Body)
	if decodeErr != nil {
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	notifyErr := r.uc.Notify(ctx, n)
	if notifyErr != nil {
		return c.String(http.StatusServiceUnavailable, notifyErr.Error())
	}

	c.Response().WriteHeader(http.StatusCreated)
	return nil
}
