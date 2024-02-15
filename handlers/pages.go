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

func HandleAlgorithmPage(c echo.Context) error {
    algorithm := c.Param("algorithm")

    return render(c, pages.Algorithm(algorithm))
}

func HandleRegisterPage(c echo.Context) error {
    return render(c, pages.Register())
}

func HandleLoginPage(c echo.Context) error {
    return render(c, pages.Login())
}

func HandleNotFoundPage(c echo.Context) error {
    return render(c, pages.NotFound())
}
