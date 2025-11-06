package main

import (
	"auction/internal/config"
	"auction/internal/connect"
	"auction/internal/container"
	"auction/internal/routes"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func main() {
	loadLocalEnv()

	cfg, err := config.LoadConfig()
	logger := setUpLogger(cfg)
	if err != nil {
		logger.Error("failed to load the config file", "error", err)
		os.Exit(1)
	}
	// _, err = connect.MongoDbConnect(cfg.MongodbUrl, cfg.MongodbPassword)
	// if err != nil {
	// 	logger.Error("failed to connect to mongodb", "error", err)
	// 	os.Exit(1)
	// }
	// logger.Info("mongodb connected successfully")

	_, err = connect.ConnectSupabase(cfg.SupbaseUrl, cfg.SupabaseToken)
	if err != nil {
		logger.Error("failed to connect to supabase servers", "error", err)
		os.Exit(1)
	}
	logger.Info("connected to supabase servers")

	appContainer, err := container.NewContainer(logger)
	if err != nil {
		logger.Error("failed to initiate new Container", "error", err)
		os.Exit(1)
	}
	router := routes.SetupRoutes(*cfg, appContainer)
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	go func() {
		logger.Info("starting server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to connect to the http server")
			os.Exit(1)
		}
	}()

	// we wait for server to shutdown gracefully
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("server shutting down", "error", err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}

	err = connect.DisconnectMongodb()
	if err != nil {
		logger.Error("failed to shutdown mongdb server", "error", err)
		os.Exit(1)
	}

	if err := connect.DisconnectSupabase(); err != nil {
		logger.Error("failed to disconnect supabase", "error", err)
	}
}

func loadLocalEnv() {
	env := strings.ToLower(strings.TrimSpace(os.Getenv("ENVIRONMENT")))
	if env == "production" {
		return
	}

	if _, err := os.Stat(".env.local"); err == nil {
		if err := godotenv.Load(".env.local"); err != nil {
			slog.Warn("Failed to load .env.local", "error", err)
		}
	}
}

func setUpLogger(cfg *config.Config) *slog.Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: cfg.IsDevelopment(),
	}
	if cfg.IsProduction() {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelDebug,
		})
	}
	return slog.New(handler)
}
