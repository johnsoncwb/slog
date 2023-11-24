package main

import (
	"log/slog"
	"os"

	"github.com/johnsoncwb/slog/slog-9/models"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	user := &models.User{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	userGroupLogger := &models.UserGroupLogger{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	userLogger := &models.UserLogger{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	slog.Info("user info", "user", user)
	slog.Info("user info", "user", userLogger)
	slog.Info("user info", "user", userGroupLogger)
}
