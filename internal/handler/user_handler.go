package handler

import (
	"go-backend/internal/model"
	"go-backend/pkg/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user in the database.
func CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if result := config.DB.Create(&u); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusCreated, u)
}

// GetUser fetches a user by ID.
func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	if result := config.DB.First(&user, id); result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, user)
}
