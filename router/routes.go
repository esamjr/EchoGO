package router

import (
	"echo/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()
	g := e.Group("/admin")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Admin Route
	g.Use(middleware.BasicAuth(validateUser))
	g.GET("/", mainAdmin)

	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.POST("/register", controllers.CreateUser)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome, Admin")
}

func validateUser(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
