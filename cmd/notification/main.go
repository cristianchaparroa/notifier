package main

import (
	"fmt"
	"notifier/config"
	"notifier/internal/service"
)

func main() {
	cfg, err := config.NewConfiguration()
	if err != nil {
		fmt.Println("Config error: %s", err)
	}

	service.Run(cfg)
}
