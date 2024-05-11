package bootstrap

import (
	"os"

	"log/slog"
)

var Log *slog.Logger

func InitLog() {
	Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
