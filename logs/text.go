package logs

import (
	"log/slog"
	"os"
)

func Text() *slog.Logger {
    return slog.New(
        slog.NewTextHandler(
            os.Stdout,
            &slog.HandlerOptions{ AddSource: false },
        ),
    )
}
