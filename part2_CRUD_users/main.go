package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validation:"numeric,required"`
	Age  int    `json:"age"`
}

var arrUser []User

func GetUsers(context echo.Context) error {
	return context.JSON(http.StatusOK, arrUser)
}

func CreateUsers(context echo.Context) error {
	var params User
	err := context.Bind(&params)
	if err != nil {
		return err
	}

	arrUser = append(arrUser, params)
	return context.JSON(http.StatusCreated, arrUser)
}

func GetUserDetail(context echo.Context) error {
	id := context.Param("id")
	for _, user := range arrUser {
		if strconv.Itoa(user.ID) == id {
			return context.JSON(http.StatusOK, user)
		}
	}
	return context.JSON(http.StatusNotFound, "Not Found")
}

// Make delete function
func DeleteUser(context echo.Context) error {

	id := context.Param("id")
	for i, buku := range arrUser {
		if strconv.Itoa(buku.ID) == id {
			arrUser = append(arrUser[:i], arrUser[i+1:]...)
			return context.JSON(http.StatusOK, buku)
		}
	}
	return context.JSON(http.StatusNotFound, "Not Found")
}

func main() {
	e := echo.New()
	arrUser = []User{
		{ID: 1, Name: "Steven", Age: 15},
	}

	// Routing
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUserDetail)
	e.POST("/users", CreateUsers)
	e.DELETE("/users/:id", DeleteUser)

	e.Start(":8000")
}
