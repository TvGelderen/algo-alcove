package handlers

import (
	"github.com/TvGelderen/algo-alcove/types"
	"github.com/TvGelderen/algo-alcove/view/pages"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
    return render(c, pages.Home())
}

func HandleRegisterPage(c echo.Context) error {
    return render(c, pages.Register(types.RegisterParams{}, types.RegisterErrors{}))
}

func HandleLoginPage(c echo.Context) error {
    return render(c, pages.Login(types.LoginParams{}, types.LoginErrors{}))
}

func HandleNotFoundPage(c echo.Context) error {
    return render(c, pages.NotFound())
}
