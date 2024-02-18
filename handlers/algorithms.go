package handlers

import (
	"net/http"
	"strconv"

	"github.com/TvGelderen/algo-alcove/models"
	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/TvGelderen/algo-alcove/view/pages"
	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) HandleAlgorithmsPage(c echo.Context) error {
	dbAlgorithms, err := h.DB.GetAlgorithmNames(c.Request().Context())
	if err != nil {
		return c.HTML(http.StatusBadRequest, "Something went wrong fetching algortihm data.")
	}

	var algorithms []models.AlgorithmName

	for _, algorithm := range dbAlgorithms {
		algorithms = append(algorithms, models.ToAlgorithmName(algorithm))
	}

	return render(c, pages.Algorithms(algorithms))
}

func HandleAlgorithmsAbout(c echo.Context) error {
	return render(c, algorithms.About())
}

func (h *DefaultHandler) HandleAlgorithm(c echo.Context) error {
	algorithmType := c.Param("type")
	algorithmName := c.Param("algorithm")

	if algorithmType == "sorting" && algorithmName == "" {
		return render(c, algorithms.Sorting())
	}
	if algorithmType == "pathfinding" && algorithmName == "" {
		return render(c, algorithms.PathFinding())
	}

	algorithm, err := h.DB.GetAlgorithmByTextId(c.Request().Context(), algorithmName)
	if err != nil {
		return render(c, pages.NotFound())
	}

	return render(c, algorithms.Algorithm(models.ToAlgorithm(algorithm)))
}

func (h *DefaultHandler) HandleGetAlogritmCode(c echo.Context) error {
	algorithmIdParam := c.Param("algorithmId")
	algorithmId, err := strconv.ParseInt(algorithmIdParam, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "Unable to parse algorithm id")
	}

	codeSlice, err := h.DB.GetAlgorithmCodeByAgorithmId(c.Request().Context(), int32(algorithmId))
	if err != nil {
		return c.String(http.StatusBadRequest, "Unable to get code for algorithm")
	}

	var codeModels []models.AlgorithmCode
	for _, code := range codeSlice {
		codeModels = append(codeModels, models.ToAlgorithmCode(code))
	}

	return render(c, algorithms.ShowCode(codeModels))
}
