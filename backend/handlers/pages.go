package handlers

import (
	"github.com/TvGelderen/algo-alcove/backend/view/pages"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
    return pages.Home().Render(c.Request().Context(), c.Response())
}
