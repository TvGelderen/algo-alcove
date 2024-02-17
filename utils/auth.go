package utils

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.RegisteredClaims
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return bytes, err
}

func CheckPasswordWithHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func CreateToken(id uuid.UUID, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Id:    id,
		Name:  name,
		Email: email,
	})

	return token.SignedString([]byte(GetHmacKey()))
}

func SetToken(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func RemoveToken(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func GetToken(r *http.Request) (string, error) {
    cookie, err := r.Cookie("AccessToken")
    if err != nil {
        return "", err
    }

    return cookie.Value, nil
}

func ParseToken(token string) (*jwt.Token, error) {
    parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(GetHmacKey()), nil
    })

    return parsedToken, err
}

func GetIdFromToken(token string) (uuid.UUID, error) {
    parsedToken, err := ParseToken(token)
    if err != nil {
        return uuid.UUID{}, err
    }

    return parsedToken.Claims.(*CustomClaims).Id, nil
}

func GetHmacKey() string {
	key := os.Getenv("HMAC_KEY")
	if key == "" {
		log.Fatal("HMAC Secret key is missing.")
	}

	return key
}
