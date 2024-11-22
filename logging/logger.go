package logging

import (
	"log/slog"
	"os"
)

func InitLogging() {
	file, _ := os.OpenFile("../logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	handler := slog.Handler(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
