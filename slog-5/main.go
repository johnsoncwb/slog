package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

/*
A criação de log "filhos", que podem ser passados para um função ou método especificos pode ajudar na centralização de informações, uma vez que o log irá levar com ele os dados setados previamente.
*/

func main() {
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	child := slog.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)

	child.Info("test message")
	/*
		{
		    "time": "2023-11-24T00:03:54.553166-03:00",
		    "level": "INFO",
		    "msg": "test message",
		    "program_info": {
		        "pid": 5859,
		        "go_version": "go1.21.4"
		    }
		}
	*/

	child.Debug("teste message 2")
	/*
			{
		    "time": "2023-11-24T00:07:06.578573-03:00",
		    "level": "DEBUG",
		    "msg": "teste message 2",
		    "program_info": {
		        "pid": 6026,
		        "go_version": "go1.21.4"
		    }
		}
	*/
}
