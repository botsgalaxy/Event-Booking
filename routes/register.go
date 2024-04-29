package routes

import (
	"net/http"
	"strconv"

	"github.com/botsgalaxy/Event-Booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
		return
	}
	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "could not register with the event",
			})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration Succesful!!!",
	})

}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
		return
	}
	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "registration canceled succesfully",
	})

}
