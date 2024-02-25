package main

import (
	"fmt"
	"notifier/config"
	"notifier/internal/service"
)

func main() {
	cfg, err := config.NewConfiguration()
	if err != nil {
		fmt.Printf("Config error: %s\n", err)
	}

	service.Run(cfg)
}
