package controllers

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17,lte=35"`
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "jinsak@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Email Not Valid",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Success",
	})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:   "Jin Sakai",
		Email:  "JinSak@gmail.com",
		Alamat: "Depok",
		Umur:   17,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Success",
	})
}
