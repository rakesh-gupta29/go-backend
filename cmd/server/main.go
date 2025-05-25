package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/portal/internal/config"
	"github.com/rakesh-gupta29/portal/internal/transport/http/router"
)

func main() {
	err := config.LoadAllConfigs()
	if err != nil {
		// TODO: use logger package
		fmt.Println("Error loading config: ", err)
	}

	fmt.Println("Config loaded successfully", config.AppConfig)

	// TODO: setup logger, swagger and other things

	app := fiber.New(config.FiberConfig)

	router.MountAllRoutes(app)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// get the meaning of the code
	go func() {
		<-sigCh
		fmt.Println("Received shutdown signal, shutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port)

	if err := app.Listen(serverAddr); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
