package service

import (
	"fmt"
	"notifier/config"
	v1 "notifier/internal/controller/http/v1"
	"notifier/internal/usecase"
	"notifier/pkg/httpserver"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) {
	em := usecase.NewEmailManager()
	uc := usecase.NewNotificationUseCase(em)
	nc := v1.NewNotificationController(uc)

	e := echo.New()
	e.POST("/notifications", nc.SendNotifications)

	httpServer := httpserver.New(e)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-interrupt:
		fmt.Println("service - Run - signal: " + i.String())
	case notifyErr := <-httpServer.Notify():
		fmt.Println("service - Run - httpServer.Notify: %w", notifyErr)
		shutdownErr := httpServer.Shutdown()
		if shutdownErr != nil {
			fmt.Println("service - Run - httpServer.Shutdown: %w", shutdownErr)
		}
	}

	shutdownErr := httpServer.Shutdown()
	if shutdownErr != nil {
		fmt.Println("service - Run - httpServer.Shutdown: %w", shutdownErr)
	}
}
