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
	BookRequest.Time = "time.Now().String()"
	BookRequest.UserID = 12                 //gin.PostForm()
	BookRequest.PickUp.Latitude = 21.222212 //gin.PostForm()
	BookRequest.PickUp.Longitude = 21.3223  //gin.PostForm()
	BookRequest.Drop.Latitude = 21.212      //gin.PostForm()
	BookRequest.Drop.Longitude = 212.22     //gin.PostForm()
	_, err := handlersImpl.svc.BookCab(BookRequest)

	if err != nil {
		g.JSON(http.StatusOK, gin.H{"message": "DONE", "status": http.StatusOK})
	} else {
		g.JSON(http.StatusOK, gin.H{"message": "ERROR", "status": http.StatusOK})
	}

}
