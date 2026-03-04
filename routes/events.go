package routes

import (
	"net/http"
	"strconv"

	"github.com/bits-and-atoms/Go_REST_API/model"
	"github.com/bits-and-atoms/Go_REST_API/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not get events",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}
func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	event, err := model.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event or event not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
func createEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized login first",
		})
		return
	}
	err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}
	var event model.Event
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse request, make sure you pass all necessary fields"})
		return
	}
	event.UserID = 1
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not get events",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	_, err = model.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "no such event",
		})
	}
	var updatedEvent model.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data",
		})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update event",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "event updated successfully",
	})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	event, err := model.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "no such event",
		})
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to delete event",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "event deleted successfully",
	})
}
