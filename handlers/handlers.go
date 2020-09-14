package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	service "github.com/Omkeshs/carbooking-backend/service"

	"github.com/Omkeshs/carbooking-backend/models"

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

//User can bookcab
func (handlersImpl HandlersImpl) BookCab(c *gin.Context) {
	BookRequest := models.CabBookRequest{}
	c.BindJSON(&BookRequest)
	fmt.Println(BookRequest)
	res, err := handlersImpl.svc.BookCab(BookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "FAILD", "status": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}

}

//user can ride history
func (handlersImpl HandlersImpl) RideHistory(c *gin.Context) {

	UserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UserId"})
	}
	res, err := handlersImpl.svc.RideHistory(UserId)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "FAILD", "status": err})
	} else {
		c.JSON(http.StatusOK, res)
	}

}
