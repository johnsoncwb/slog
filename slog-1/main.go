package main

import "log/slog"

/*
Estrutura básica de logs com slog, pode-se notar que o package utiliza como base o nível Info, por isso a menssagem de Debug é suprimida ao rodar o programa
*/

func main() {
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}
