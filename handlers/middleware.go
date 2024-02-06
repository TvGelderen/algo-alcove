package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func validAuthToken(r *http.Request) (uuid.UUID, error) {
    token, err := getToken(r)
    if err != nil {
        return uuid.UUID{}, err
    }

    id, err := getIdFromToken(token)
    if err != nil {
        return uuid.UUID{}, err
    }

    return id, nil
}

func DefaultMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        sessionId := uuid.New()
        c.Set("session-id", sessionId)

        id, err := validAuthToken(c.Request())
        if err == nil {
            c.Set("user-id", id)
        }

        fmt.Printf("New session: %v\n", sessionId)
        
        return next(c)
    }
}

func AuthorizePage(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Get("user-id")
        if id == nil {
            return c.Redirect(302, "/login")
        }

        return next(c)
    }
}
