package main

import (
	"fmt"
	"os"

	"github.com/TvGelderen/algora/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT is missing, defaulting to 3000")
		port = ":3000"
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.Static("/assets/*", "assets")

	e.GET("/", handlers.HandleHomePage)

	e.GET("/algorithms", handlers.HandleAlgorithmsAbout)
	e.GET("/algorithms/sorting", handlers.HandleAlgorithmsSortingAbout)
	e.GET("/algorithms/pathfinding", handlers.HandleAlgorithmsPathFindingAbout)
	e.GET("/algorithms/sorting/bubble-sort", handlers.HandleBubbleSort)
	e.GET("/algorithms/sorting/insertion-sort", handlers.HandleInsertionSort)
	e.GET("/algorithms/sorting/selection-sort", handlers.HandleSelectionSort)
	e.GET("/algorithms/sorting/merge-sort", handlers.HandleMergeSort)
	e.GET("/algorithms/sorting/quick-sort", handlers.HandleQuickSort)
	e.GET("/algorithms/pathfinding/breadth-first-search", handlers.HandleBreadthFirstSearch)
	e.GET("/algorithms/pathfinding/depth-first-search", handlers.HandleDepthFirstSearch)
	e.GET("/algorithms/pathfinding/dijkstras-algorithm", handlers.HandleDijkstrasAlgorithm)
	e.GET("/algorithms/pathfinding/a-star-algorithm", handlers.HandleAStarAlgorithm)

	e.GET("/*", handlers.HandleNotFoundPage)

	fmt.Printf("Starting server on %s", port)

	e.Start(port)
}
