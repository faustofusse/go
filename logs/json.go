package logs

import (
	"log/slog"
	"os"
)

func Json() *slog.Logger {
    return slog.New(
        slog.NewJSONHandler(
            os.Stdout,
            &slog.HandlerOptions{ AddSource: false },
        ),
    )
}
