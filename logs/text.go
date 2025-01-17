package logs

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

const timeFormat = "15:04:02"

func Text() *slog.Logger {
    options := slog.HandlerOptions{}
    options.Level = slog.LevelDebug
    options.AddSource = false
    options.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
        return slog.Any("", nil)
    }
    handler := TextHandler{
        buffer: new(bytes.Buffer),
        mutex: new(sync.Mutex),
        wrapped: slog.NewTextHandler(os.Stdout, &options),
    }
    return slog.New(handler)
}

type TextHandler struct {
	wrapped slog.Handler
	buffer *bytes.Buffer
	mutex *sync.Mutex
}

func (h TextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.wrapped.Enabled(ctx, level)
}

func (h TextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TextHandler{ wrapped: h.wrapped.WithAttrs(attrs), buffer: h.buffer, mutex: h.mutex }
}

func (h TextHandler) WithGroup(name string) slog.Handler {
	return &TextHandler{ wrapped: h.wrapped.WithGroup(name), buffer: h.buffer, mutex: h.mutex }
}

func levelString(level slog.Level) string {
    switch level {
        case slog.LevelDebug: return "D"
        case slog.LevelInfo: return "I"
        case slog.LevelWarn: return "W"
        case slog.LevelError: return "E"
        default: return "?"
    }
}

func levelColor(level slog.Level) int {
    switch level {
        case slog.LevelDebug: return darkGray
        case slog.LevelInfo: return cyan
        case slog.LevelWarn: return lightYellow
        case slog.LevelError: return lightRed
        default: return white
    }
}

func rightPad(str string, width int) string {
    padded := str
    for len(padded) < width {
        padded += " "
    }
    return padded
}

func (h TextHandler) Handle(ctx context.Context, r slog.Record) error {
    fmt.Print(
        colorize(darkGray, r.Time.Format(timeFormat)),
        " ",
        colorize(levelColor(r.Level), levelString(r.Level)),
        " ",
        colorize(levelColor(r.Level), "│"),
        " ",
    )

    attrs := map[string]string{}

    r.Attrs(func(a slog.Attr) bool {
        attrs[a.Key] = a.Value.String()
        return true
    })

    service := attrs["service"]

    if service == "http" {
        if len(attrs["method"]) > 4 {
            attrs["method"] = attrs["method"][:3]
        }
        fmt.Print(
            colorize(white, fmt.Sprintf("%-4s", attrs["method"])),
            " ",
            colorize(white, attrs["status"]),
            " ",
            colorize(white, attrs["path"]),
            " ",
            colorize(darkGray, attrs["latency_human"]),
        )
    } else {
        if service != "" {
            fmt.Print(colorize(lightGray, service), ": ")
        }
        fmt.Print(colorize(white, r.Message))
        for key, value := range attrs {
            if key != "service" {
                fmt.Print(
                    " ",
                    colorize(lightGray, key),
                    colorize(lightGray, "="),
                    colorize(white, value),
                )
            }
        }
    }

    fmt.Print("\n")

    return nil
}
