package controllers

import (
	"net/http"

	"github.com/myproject/mysqlRestAPI/models"

	"github.com/labstack/echo"
)

func getAllUsers(c echo.Context) error {
	result := models.getAllUsers()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
