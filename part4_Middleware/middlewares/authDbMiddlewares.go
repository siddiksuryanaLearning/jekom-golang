package middlewares

import (
	"github.com/labstack/echo"
	"jekom-golang/part4_Middleware/config"
	"jekom-golang/part4_Middleware/models"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
