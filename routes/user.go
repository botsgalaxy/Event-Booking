package routes

import (
	"fmt"
	"net/http"

	"github.com/botsgalaxy/Event-Booking/models"
	"github.com/botsgalaxy/Event-Booking/utils"
	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!!!",
	})

}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
	}
	err = user.ValidateCredentials()
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "could not authenticate user",
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful!",
		"token":   token,
	})

}
