package handlers

import (
	"github.com/TvGelderen/algo-alcove/view/pages"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
    return render(c, pages.Home())
}

func HandleAlgorithmsPage(c echo.Context) error {
    return render(c, pages.Algorithms())
}
