package logs

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func bytesIn(req *http.Request) string {
    cl := req.Header.Get(echo.HeaderContentLength)
    if cl == "" {
        cl = "0"
    }
    return cl
}

// Logger returns a middleware that logs HTTP requests.
func Echo(logger *slog.Logger) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(ctx echo.Context) error {
			req := ctx.Request()
			res := ctx.Response()
			start := time.Now()
            err := next(ctx)
			stop := time.Now()
            if err != nil {
				ctx.Error(err)
                logger.Error(err.Error())
			}
            logger.Info(
                "request",
                "service",
                "http",
                "status", res.Status,
                "method", req.Method,
                "path", req.URL.Path,
				"time_unix_nano", start.UnixNano(),
                "remote_ip", ctx.RealIP(),
				// "host", req.Host,
                "bytes_in", bytesIn(ctx.Request()),
                "bytes_out", res.Size,
                "latency_nano", stop.Sub(start).Nanoseconds(),
                "latency_human", stop.Sub(start).String(),
                "user_agent", req.UserAgent(),
				// "error", reqError(err),
            )
            return nil
        }
    }
}
