package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TvGelderen/algo-alcove/handlers"
	"github.com/TvGelderen/algo-alcove/seeds"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	args := flag.Args()

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

	if len(args) >= 1 && args[0] == "seed" {
		seeds.Seed(dbConnectionString)

		return
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

	e.GET("/*", handlers.HandleBasePage, h.DefaultPageMiddleware)

	e.GET("/pages/", handlers.HandleHomePage)
	e.GET("/pages/register", handlers.HandleRegisterPage)
	e.GET("/pages/login", handlers.HandleLoginPage)

	e.PUT("/api/register", h.HandleRegister)
	e.POST("/api/login", h.HandleLogin)
	e.GET("/api/logout", h.HandleLogout)

    addAlgorithmEndpoints(e, h)

	e.GET("/pages/*", handlers.HandleNotFoundPage)
}

func addAlgorithmEndpoints(e *echo.Echo, h *handlers.DefaultHandler) {
	e.GET("/pages/algorithms", h.HandleAlgorithmsPage)
    e.GET("/pages/algorithms/*", h.HandleAlgorithmsPage)

	e.GET("/api/algorithms", handlers.HandleAlgorithmsAbout)
    e.GET("/api/algorithms/:type", h.HandleAlgorithm)
    e.GET("/api/algorithms/:type/:algorithm", h.HandleAlgorithm)
	e.GET("/api/algorithms/:algorithmId/code", h.HandleGetAlogrithmCode)
}
