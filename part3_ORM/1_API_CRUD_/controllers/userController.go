package controllers

import (
	"net/http"

	"alta.id/go-skeleton/lib/database"
	"alta.id/go-skeleton/models"
	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	result := database.DeleteUser(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan server, gagal medapatkan data buku.",
		})
	}
}
