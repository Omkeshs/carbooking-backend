package endpoints

import (
	handler "github.com/Omkeshs/carbooking-backend/handlers"

	"github.com/gin-gonic/gin"
)

//NewRoute All Application Routes Are defiend Here
func NewRoute(router *gin.Engine, handler *handler.HandlersImpl) {
	router.POST("/bookcab", handler.BookCab)
	router.GET("/history/:id", handler.RideHistory)
}
