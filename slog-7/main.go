package main

import (
	"context"
	"log/slog"
	"os"
)

/*
podemos criar níveis de log personalizados, uma vez que os níveis built-in são:
DEBUG (-4), INFO (0), WARN (4), and ERROR (8)

*/

const (
	LevelTrace  = slog.Level(-8)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelTrace:  "TRACE",
	LevelNotice: "NOTICE",
	LevelFatal:  "FATAL",
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: LevelTrace,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}))

	slog.SetDefault(logger)

	ctx := context.Background()

	slog.Log(ctx, LevelTrace, "Trace message")
	slog.Log(ctx, LevelNotice, "Trace message")
	slog.Log(ctx, LevelFatal, "Trace message")
}
