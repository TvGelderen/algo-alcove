package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
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

	fs := http.FileServer(http.Dir("assets"))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    e.GET("/", handlers.HandleHomePage, h.DefaultPageMiddleware)

    e.GET("/algorithms", handlers.HandleAlgorithmsPage, h.DefaultPageMiddleware)
    e.GET("/algorithms/*", handlers.HandleAlgorithmsPage, h.DefaultPageMiddleware)

    e.GET("/api/algorithms", handlers.HandleAlgorithmsAbout)
    e.GET("/api/algorithms/:type", handlers.HandleAlgorithm)
    e.GET("/api/algorithms/:type/:algorithm", handlers.HandleAlgorithm)

    e.GET("/register", handlers.HandleRegisterPage, h.DefaultPageMiddleware)
    e.GET("/login", handlers.HandleLoginPage, h.DefaultPageMiddleware)

    e.PUT("/api/register", h.HandleRegister)
    e.POST("/api/login", h.HandleLogin)
    e.GET("/api/logout", h.HandleLogout)

    e.GET("/*", handlers.HandleNotFoundPage)

    e.Start(":" + port)
}
