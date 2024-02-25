package v1

import (
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
	return nil
}
