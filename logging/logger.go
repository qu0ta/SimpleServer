// Package logging contains the entry point of the program and the main api handles
//
// This package includes logger
package logging

import (
	"log/slog"
	"os"
)

// InitLogging initializes the logging system for the application.
//
// It creates a file named "gin.log" in the "../logs" directory and opens it in append mode.
// It then creates a new JSON handler with the specified log level and source information.
// The handler is used to create a new logger, which is set as the default logger.
//
// Parameters:
// None.
//
// Returns:
// None.
func InitLogging() {
	file, _ := os.OpenFile("../logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	handler := slog.Handler(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
