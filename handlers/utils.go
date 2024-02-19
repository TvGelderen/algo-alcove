package handlers

import (
	"database/sql"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	DB *database.Queries
}

func New(db *sql.DB) *DefaultHandler {
	return &DefaultHandler{
		DB: database.New(db),
	}
}

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
