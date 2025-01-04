package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func redirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("Hx-Request")) > 0 {
		c.Response().Writer.Header().Set("Hx-Redirect", to)
		return nil
	}

	return c.Redirect(302, to)
}
