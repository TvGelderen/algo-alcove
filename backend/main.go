package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TvGelderen/algo-alcove/backend/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    godotenv.Load(".env")

	port := os.Getenv("AA_PORT")
	if port == "" {
		port = "8080"
        fmt.Print("PORT is missing, defaulting to 8080\n")
	}

	dbConnectionString := os.Getenv("AA_DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("No database connection string found.")
	}

	// connection, err := sql.Open("postgres", dbConnectionString)
	// if err != nil {
	// 	log.Fatal("Unable to establish connection with database: ", err)
	// }

    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	fs := http.FileServer(http.Dir("assets"))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    e.GET("/", handlers.HandleHomePage)

    e.Start(":" + port)
}
