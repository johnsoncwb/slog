package main

import (
	"context"
	"log/slog"
	"os"
)

/*
temos a possibilidade de criar grupos dentro do log.
*/

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties",
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)

	/*
			{
		    "time": "2023-11-23T23:59:04.172717-03:00",
		    "level": "INFO",
		    "msg": "image uploaded",
		    "id": 23123,
		    "properties": {
		        "width": 4000,
		        "height": 3000,
		        "format": "jpeg"
		    }
		}
	*/
}
