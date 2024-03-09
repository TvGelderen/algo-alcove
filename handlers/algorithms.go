package handlers

import (
	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/labstack/echo/v4"
)

func HandleAlgorithmsAbout(c echo.Context) error {
	return render(c, algorithms.About())
}

func HandleAlgorithmsSortingAbout(c echo.Context) error {
	return render(c, algorithms.Sorting())
}

func HandleAlgorithmsPathFindingAbout(c echo.Context) error {
	return render(c, algorithms.PathFinding())
}

func HandleBubbleSort(c echo.Context) error {
	return render(c, algorithms.BubbleSort())
}
