package logs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Echo() echo.MiddlewareFunc {
    return middleware.Logger()
}
