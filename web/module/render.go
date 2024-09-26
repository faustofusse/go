package module

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func buildUrl(parts []string, module string) string {
    result := "/" + module
    for i, part := range parts {
        if i == 0 && strings.Contains(part, module) {
            continue
        }
        if len(part) == 0 {
            continue
        }
        if part[0] != '/' {
            result += "/"
        }
        result += part
    }
    return result
}

func (handler *Handler) Render(ctx echo.Context, component templ.Component, url ...string) error {
    if len(url) > 0 {
        ctx.Response().Header().Set("Hx-Push-Url", buildUrl(url, handler.ModuleUrl))
    }
    return component.Render(ctx.Request().Context(), ctx.Response())
}
