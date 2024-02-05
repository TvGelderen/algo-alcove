package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/TvGelderen/algo-alcove/database"
)

type CustomClaims struct {
    Id uuid.UUID `json:"id"`
    Email string `json:"email"`
    Name string `json:"name"`
    jwt.RegisteredClaims
}

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

func (h *DefaultHandler) HandleLogin(c echo.Context) error {
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

    user, err := h.DB.GetUserByEmail(c.Request().Context(), params.Email)
    if err != nil {
		return c.HTML(http.StatusUnauthorized, "Wrong email or password.")
    }

    validPassword := checkPasswordWithHash(params.Password, user.PasswordHash)
    if !validPassword {
		return c.HTML(http.StatusUnauthorized, "Wrong email or password.")
    }

    token, err := createToken(user.ID, user.Username, user.Email)
    if err != nil {
		return c.HTML(http.StatusBadRequest, "Something went wrong.")
    }

    setToken(c.Response().Writer, token)

    c.Response().Writer.Header().Set("Hx-Redirect", "/")
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

func createToken(id uuid.UUID, name, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
        Id: id,
        Name: name,
        Email: email,
    })

    return token.SignedString([]byte(getHmacKey()))
}

func setToken(w http.ResponseWriter, token string) {
    cookie := http.Cookie{
        Name: "AccessToken",
        Value: token,
        MaxAge: 3600,
        Path: "/",
        HttpOnly: true,
    }

    http.SetCookie(w, &cookie)
}

func getHmacKey() string {
    key := os.Getenv("HMAC_KEY")
    if key == "" {
        log.Fatal("HMAC Secret key is missing.")
    }

    return key
}
