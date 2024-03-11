package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/TvGelderen/algo-alcove/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Print("PORT is missing, defaulting to 3000\n")
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("No database connection string found.")
	}

	connection, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}

	h := handlers.New(connection)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Use(handlers.DefaultMiddleWare)

    addRoutes(e, h)

	e.Start(":" + port)
}

func addRoutes(e *echo.Echo, h *handlers.DefaultHandler) {
	e.Static("/assets/*", "assets")

	e.GET("/", handlers.HandleHomePage, h.DefaultPageMiddleware)

	e.GET("/register", handlers.HandleRegisterPage)
	e.GET("/login", handlers.HandleLoginPage)

	e.PUT("/api/register", h.HandleRegister)
	e.POST("/api/login", h.HandleLogin)
	e.GET("/api/logout", h.HandleLogout)

    addAlgorithmEndpoints(e)

	e.GET("/*", handlers.HandleNotFoundPage)
}

func addAlgorithmEndpoints(e *echo.Echo) {
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
}
