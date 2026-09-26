package main

import (
	"booking/internal/app"
	"log/slog"
	"os"
)

func main() {
	application := app.NewApp()
	if err := application.Run(); err != nil {
		slog.Error(
			"Application Error",
			"error", err,
		)
		os.Exit(1)
	}
}
