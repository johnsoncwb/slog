package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/johnsoncwb/slog/slog-8/logger"
)

func main() {
	jsonLogger := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	handler := logger.Handler{
		Handler: jsonLogger,
	}

	slog.SetDefault(slog.New(handler))

	ctx := context.Background()

	ctx = context.WithValue(ctx, logger.ReqID, "12345")

	slog.DebugContext(ctx, "message_with_context", slog.String("url", "localhost:8080"))

	/*
			{
		    "time": "2023-11-24T00:34:07.950233-03:00",
		    "level": "DEBUG",
		    "msg": "message_with_context",
		    "url": "localhost:8080",
		    "request_id": "12345"
		}
	*/
}
