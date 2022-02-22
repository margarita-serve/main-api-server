package main

import (
	_ "git.k3.acornsoft.io/msit-auto-ml/koreserv/docs"
	models_public_http "git.k3.acornsoft.io/msit-auto-ml/koreserv/pkg/models/interfaces/public/http"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title KoreServ Swagger API
// @version 1.0
// @host localhost:30000
// @BasePath /api/v1
func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	models_public_http.AddRoutes(e)

	e.Logger.Fatal(e.Start(":30000"))
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user/{name} [get]
// func getUser(c echo.Context) error {
// 	user := new(User)
// 	// 중략
// 	return c.JSONPretty(http.StatusOK, *user, "  ")
// }

// // @Summary Create user
// // @Description Create new user
// // @Accept json
// // @Produce json
// // @Param userBody body User true "User Info Body"
// // @Success 200 {object} User
// // @Router /user [post]
// func createUser(c echo.Context) error {
// 	user := new(User)
// 	// 중략
// 	return c.JSONPretty(http.StatusOK, *user, "  ")
// }
