package repositories

import (
	"fmt"
	"std/omkesh/carBooking-Backend/models"

	"github.com/jinzhu/gorm"
)

//Repository implimets all methods in Repository
type Repository interface {
	BookCab(models.CabBookRequest) (*models.CabBookResponce, error)
}

//RepositoryImpl **
type RepositoryImpl struct {
	dbConn *gorm.DB
}

//NewRepositoryImpl inject dependancies of DataStore
func NewRepositoryImpl(dbConn *gorm.DB) Repository {
	return &RepositoryImpl{dbConn: dbConn}
}

func (repositoryImpl RepositoryImpl) BookCab(req models.CabBookRequest) (*models.CabBookResponce, error) {
	dbConn := repositoryImpl.dbConn

	fmt.Println("STEP 5 : FROM REPO", req, dbConn)
	resp := models.CabBookResponce{}
	resp.BookingID = 12
	resp.CabNo = "MH14DZ7416"
	resp.DriverMob = "83989283"
	resp.DriverName = "om"
	return &resp, nil
}
