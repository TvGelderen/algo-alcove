package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	fs := http.FileServer(http.Dir("assets"))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    e.GET("/", handlers.HandleHomePage, handlers.AuthorizePage)
    e.GET("/algorithms", handlers.HandleAlgorithmsPage)
    e.GET("/register", handlers.HandleRegisterPage)
    e.GET("/login", handlers.HandleLoginPage)

    e.PUT("/api/register", h.HandleRegister)
    e.POST("/api/login", h.HandleLogin)

    e.Start(":" + port)
}
