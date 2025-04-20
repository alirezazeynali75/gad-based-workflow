package main

import (
    "log/slog"
    "os"

    "github.com/alirezazeynali75/gad-based-workflow/api"
    "github.com/alirezazeynali75/gad-based-workflow/internal/configs"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configurations
    cfg, err := configs.Configure()
    if err != nil {
        slog.With("err", err.Error()).Error("failed to load configurations")
        os.Exit(1)
    }

    // Initialize logger
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    // Initialize Gin router
    router := gin.Default()

    // Initialize handlers
    handlers := api.NewHandlers(logger)

    // Register routes
    handlers.RegisterRoutes(router)

    // Start the server
    address := cfg.Http.Address + ":" + cfg.Http.Port
    logger.Info("starting server", "address", address)
    if err := router.Run(address); err != nil {
        logger.With("err", err.Error()).Error("failed to start server")
        os.Exit(1)
    }
}