package handlers

import (
	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/labstack/echo/v4"
)

func HandleAlgorithm(c echo.Context) error {
	algorithm := c.QueryParam("algorithm")

	return render(c, algorithms.Algorithm(algorithm))
}

func HandleSortingAlgorithm(c echo.Context) error {
	algorithm := c.Param("algorithm")

	if algorithm == "" {
		return render(c, algorithms.Sorting())
	}

    return render(c, algorithms.Algorithm(algorithm))
}

func HandlePathFindingAlgorithm(c echo.Context) error {
	algorithm := c.Param("algorithm")

	if algorithm == "" {
        return render(c, algorithms.PathFinding())
	}

    return render(c, algorithms.Algorithm(algorithm))
}
