package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

func (h *DefaultHandler) DefaultPageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Get("user-id")
		if id != nil {
			user, err := h.DB.GetUserById(c.Request().Context(), id.(uuid.UUID))
			if err == nil {
                ctx := c.Request().Context()
                ctx = context.WithValue(ctx, "user-email", user.Email)
                c.SetRequest(c.Request().WithContext(ctx))

                fmt.Printf("set email to: %v\n", user.Email)
            }
		}

		return next(c)
	}
}

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
