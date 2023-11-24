package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

/*
Junto dos logs filhos, é possível adicionar dados para aparecerem dentro de grupos pré criados.
*/

func main() {
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).WithGroup("program_info")

	slog.SetDefault(logger)

	child := slog.With(
		slog.Int("pid", os.Getpid()),
		slog.String("go_version", buildInfo.GoVersion),
	)
	child.Info("test_message_1", slog.String("image_id", "39ud88"))
	child.Warn("test_message_2", slog.String("available_space", "900.1 MB"))
	/*
			{
		    "time": "2023-11-24T00:13:13.950222-03:00",
		    "level": "INFO",
		    "msg": "test_message_1",
		    "program_info": {
		        "pid": 6750,
		        "go_version": "go1.21.4",
		        "image_id": "39ud88"
		    }
		}
		{
		    "time": "2023-11-24T00:13:13.950669-03:00",
		    "level": "WARN",
		    "msg": "test_message_2",
		    "program_info": {
		        "pid": 6750,
		        "go_version": "go1.21.4",
		        "available_space": "900.1 MB"
		    }
		}
	*/
}
