package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *DefaultHandler) HandleRegister(c echo.Context) error {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	err := json.NewDecoder(c.Request().Body).Decode(&params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.HTML(http.StatusBadRequest, "Something went wrong decoding.")
	}

	passwordHash, err := hashPassword(params.Password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.HTML(http.StatusBadRequest, "Something went wrong.")
	}

	err = h.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:           uuid.New(),
		Username:     "test",
		Email:        params.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if strings.Contains(err.Error(), "users_email_key") {
			return c.HTML(http.StatusUnauthorized, "That email is already taken.")
		}
		return c.HTML(http.StatusUnauthorized, "Something went wrong creating account.")
	}

	c.Response().Writer.Header().Set("Hx-Redirect", "/login")

	return nil
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return bytes, err
}

func checkPasswordWithHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
