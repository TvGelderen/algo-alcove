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
