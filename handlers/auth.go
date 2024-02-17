package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/TvGelderen/algo-alcove/types"
	"github.com/TvGelderen/algo-alcove/utils"
	"github.com/TvGelderen/algo-alcove/view/pages"
)

func (h *DefaultHandler) HandleRegister(c echo.Context) error {
	params := types.RegisterParams{}
	err := json.NewDecoder(c.Request().Body).Decode(&params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
        return render(c, pages.RegisterForm(params, types.RegisterErrors{ Other: "Something went wrong" }))
	}

    if errors, hasErrors := params.Validate(); hasErrors {
        return render(c, pages.RegisterForm(params, errors))
    }

	passwordHash, err := utils.HashPassword(params.Password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
        return render(c, pages.RegisterForm(params, types.RegisterErrors{ Other: "Something went wrong" }))
	}

	err = h.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:           uuid.New(),
		Username:     "test",
		Email:        params.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if strings.Contains(err.Error(), "users_email_key") {
            return render(c, pages.RegisterForm(params, types.RegisterErrors{ Email: "That email is already taken" }))
		}
        return render(c, pages.RegisterForm(params, types.RegisterErrors{ Other: "Something went wrong" }))
	}

	c.Response().Writer.Header().Set("Hx-Redirect", "/login")
	return nil
}

func (h *DefaultHandler) HandleLogin(c echo.Context) error {
	params := types.LoginParams{}
	err := json.NewDecoder(c.Request().Body).Decode(&params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
        return render(c, pages.LoginForm(params, types.LoginErrors{ Other: "Something went wrong" }))
	}

    if errors, hasErrors := params.Validate(); hasErrors {
        return render(c, pages.LoginForm(params, errors))
    }

	user, err := h.DB.GetUserByEmail(c.Request().Context(), params.Email)
	if err != nil {
        return render(c, pages.LoginForm(params, types.LoginErrors{ Other: "Wrong email or password" }))
	}

	validPassword := utils.CheckPasswordWithHash(params.Password, user.PasswordHash)
	if !validPassword {
        return render(c, pages.LoginForm(params, types.LoginErrors{ Other: "Wrong email or password" }))
	}

	token, err := utils.CreateToken(user.ID, user.Username, user.Email)
	if err != nil {
        return render(c, pages.LoginForm(params, types.LoginErrors{ Other: "Something went wrong" }))
	}

	utils.SetToken(c.Response().Writer, token)

	c.Response().Writer.Header().Set("Hx-Redirect", "/")
	return nil
}

func (h *DefaultHandler) HandleLogout(c echo.Context) error {
    utils.RemoveToken(c.Response().Writer)

    return c.Redirect(302, "/")
}
