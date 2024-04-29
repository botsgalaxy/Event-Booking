package routes

import (
	"net/http"
	"strconv"

	"github.com/botsgalaxy/Event-Booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	event.UserID = c.GetInt64("userId")
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return

	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})

}

func getEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event!",
		})
		return
	}
	c.JSON(http.StatusOK, event)

}

func updateEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event!",
		})
		return
	}

	userId := c.GetInt64("userId")
	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to update this event",
		})
		return

	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event data",
		})
		return
	}
	updatedEvent.ID = event.ID
	updatedEvent.Update()
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated succesfully !!!",
	})
}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event!",
		})
		return
	}
	userId := c.GetInt64("userId")
	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to delete this event",
		})
		return

	}
	event.Delete()
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted succesfully !!!",
	})

}
