package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func LoadJSONLogger() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
