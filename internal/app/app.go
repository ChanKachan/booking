package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App структура нашего приложения
type App struct {
	server *http.Server
	logger *slog.Logger
}

func NewApp() *App {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	router := gin.New()
	router.Use(gin.Recovery())

	// Todo: Временно помести endpoint здесь, позже перенесем в handlers
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
		return
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return &App{
		server: server,
		logger: logger,
	}
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	serverErr := make(chan error, 1)

	go func() {
		a.logger.Info(
			"http server started",
			"addr", a.server.Addr,
		)

		serverErr <- a.server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		a.logger.Info("shutdown started")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		// Завершает все текущие запросы, только после останавливается
		if err := a.server.Shutdown(shutdownCtx); err != nil {
			return err
		}

		a.logger.Info("application stopped")
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("http server failed: %w", err)
		}
	}

	return nil
}
