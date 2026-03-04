package routes

import (
	"net/http"
	"strconv"

	"github.com/bits-and-atoms/Go_REST_API/model"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	event_id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	user_id := ctx.GetInt64("userId")
	event, err := model.GetEventById(event_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "no such event",
		})
	}
	err = event.Register(user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user registered for event successfully"})
}
func cancelRegistration(ctx *gin.Context) {

}
