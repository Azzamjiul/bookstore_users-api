package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		// TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println(saveErr.Error())
		// TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

}

func SearchUser(c *gin.Context) {

}
