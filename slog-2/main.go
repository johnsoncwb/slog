package main

import (
	"log"
	"log/slog"
	"os"
)

/*
existe a possibilidade de se utilizar dois tipos built-in de handlers para o slog, o JSONHandler e o TextHandler, para os nossos casos onde os serviços compreender logs baseados em json, podemos focar no JSONHandler.
*/

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	log.Println("message from old logger")
}
