package handlers

import (
	"fmt"

	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/labstack/echo/v4"
)

func HandleAlgorithmsAbout(c echo.Context) error {
    return render(c, algorithms.About())
}

func HandleAlgorithm(c echo.Context) error {
	algorithmType := c.Param("type")
	algorithmName := c.Param("algorithm")

    fmt.Printf("type: %v\n", algorithmType)
    fmt.Printf("algorithmName: %v\n", algorithmName)
    
	if algorithmType == "sorting" && algorithmName == "" {
		return render(c, algorithms.Sorting())
	}
	if algorithmType == "pathfinding" && algorithmName == "" {
        return render(c, algorithms.PathFinding())
	}

	return render(c, algorithms.Algorithm(algorithmName))
}
