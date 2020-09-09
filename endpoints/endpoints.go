package endpoints

import (
	"fmt"
	handler "std/omkesh/carBooking-Backend/handlers"

	"github.com/gin-gonic/gin"
)

//NewRoute All Application Routes Are defiend Here
func NewRoute(router *gin.Engine, handler *handler.HandlersImpl) {
	fmt.Println("STEP 2 : FROM ENDPOINT")
	router.POST("/bookcab", handler.BookCab)
}
