package handlers

import (
	"github.com/TvGelderen/algo-alcove/view/algorithms"
	"github.com/TvGelderen/algo-alcove/view/pages"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
	return render(c, pages.Home())
}

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

func HandleInsertionSort(c echo.Context) error {
	return render(c, algorithms.InsertionSort())
}

func HandleSelectionSort(c echo.Context) error {
	return render(c, algorithms.SelectionSort())
}

func HandleMergeSort(c echo.Context) error {
	return render(c, algorithms.MergeSort())
}

func HandleQuickSort(c echo.Context) error {
	return render(c, algorithms.QuickSort())
}

func HandleBreadthFirstSearch(c echo.Context) error {
	return render(c, algorithms.BreadthFirstSearch())
}

func HandleDepthFirstSearch(c echo.Context) error {
	return render(c, algorithms.DepthFirstSearch())
}

func HandleDijkstrasAlgorithm(c echo.Context) error {
	return render(c, algorithms.DijkstrasAlgorithm())
}

func HandleAStarAlgorithm(c echo.Context) error {
	return render(c, algorithms.AStarAlgorithm())
}

func HandleNotFoundPage(c echo.Context) error {
	return render(c, pages.NotFound())
}
