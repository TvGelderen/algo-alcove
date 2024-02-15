package handlers

import (
	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/TvGelderen/algo-alcove/view/pages"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
    return render(c, pages.Home())
}

func HandleAlgorithmsAboutPage(c echo.Context) error {
    return render(c, algorithms.About())
}

func HandleSortingAlgorithmsPage(c echo.Context) error {
    return render(c, algorithms.Sorting())
}

func HandlePathFindinAlgorithmsPage(c echo.Context) error {
    return render(c, algorithms.PathFinding())
}

func HandleAlgorithmPage(c echo.Context) error {
    algorithm := c.Param("algorithm")

    return render(c, algorithms.Algorithm(algorithm))
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
