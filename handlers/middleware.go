package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TvGelderen/algo-alcove/models"
	"github.com/TvGelderen/algo-alcove/utils"
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
            return redirect(c, "/login")
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
                ctx = context.WithValue(ctx, "user", models.ToUser(user))
                c.SetRequest(c.Request().WithContext(ctx))
            }
		}

		return next(c)
	}
}

func validAuthToken(r *http.Request) (uuid.UUID, error) {
	token, err := utils.GetToken(r)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := utils.GetIdFromToken(token)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
