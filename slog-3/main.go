package main

import (
	"context"
	"log/slog"
	"os"
)

/*
exemplo de log com informações complementares, todas os níveis de log aceitam uma mensagem como seu primeiro parametro e um número ilimitado de key/values após.
mas deve-se tomar cuidado para não termos um erro como mostado no log de warning abaixo, onde só é setado um valor sem chave, então o slog completa o log com "!BADKEY:<valor>"
pode-se tomar uma abordagem mais cuidadosa ao criar seus parametros do log e utilizar os "strongly-typed contextual attributes".
*/

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	slog.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	slog.Warn(
		"incoming request",
		"method", "GET",
		"time_taken_ms", // the value for this key is missing
	)

	slog.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	slog.Info(
		"incoming request",
		"method", "GET",
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		"status", 200,
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

	/*
		método "seguro" para logar parametros, uma vez que só é aceito valores do tipo slog.Attr, que são compostos por uma chave e valor, este método é mais complicado que o mostrado acima, uma vez que se faz necessário o uso de um contexto (ou nil) e você deve indicar o nível do log a cada chamada
	*/

	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)
}
