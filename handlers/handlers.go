package handlers

import (
	"fmt"
	"net/http"
	service "std/omkesh/carBooking-Backend/service"

	"std/omkesh/carBooking-Backend/models"

	"github.com/gin-gonic/gin"
)

//HandlersImpl for handler Functions
type HandlersImpl struct {
	svc service.Service
}

//NewHandlerImpl inits dependancies for graphQL and Handlers
func NewHandlerImpl(service service.Service) *HandlersImpl {
	return &HandlersImpl{svc: service}
}

//bookcab
func (handlersImpl HandlersImpl) BookCab(g *gin.Context) {
	fmt.Println("STEP 3 : FROM HANDLER")

	BookRequest := models.CabBookRequest{}
	g.BindJSON(&BookRequest)
	fmt.Println(BookRequest)
	_, err := handlersImpl.svc.BookCab(BookRequest)

	if err != nil {
		g.JSON(http.StatusOK, gin.H{"message": "DONE", "status": http.StatusOK})
	} else {
		g.JSON(http.StatusOK, gin.H{"message": "ERROR", "status": http.StatusOK})
	}

}
