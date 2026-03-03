package main

import (
	"net/http"
	"strconv"

	"github.com/bits-and-atoms/Go_REST_API/db"
	"github.com/bits-and-atoms/Go_REST_API/model"
	"github.com/gin-gonic/gin"
)
func main(){
	db.InitDB()
	server := gin.Default()
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{"message":"healthy"})
	})
	server.GET("/events", getEvents)
	server.GET("/events/:id",getEvent)
	server.POST("/events",createEvent)
	server.Run(":8080")
}
func getEvents(ctx *gin.Context){
	events ,err := model.GetAllEvents()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not get events",
		})
		return
	}
	ctx.JSON(http.StatusOK,events)
}
func getEvent(ctx *gin.Context){
	id,err := strconv.ParseInt(ctx.Param("id"),10,64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"could not parse event id",
		})
		return 
	}
	event,err := model.GetEventById(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not fetch event or event not found",
		})
		return
	}
	ctx.JSON(http.StatusOK,event)
}
func createEvent(ctx *gin.Context){
	var event model.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"unable to parse request, make sure you pass all necessary fields"})
		return
	}
	event.UserID = 1
	err = event.Save()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not get events",
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{"message":"event created","event":event})
}