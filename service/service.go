package services

import (
	"carbooking-backend/models"
	"errors"

	repository "carbooking-backend/repositories"

	helpers "carbooking-backend/helpers"
)

// Service describes the service.
type Service interface {
	BookCab(models.CabBookRequest) (models.CabBookResponce, error)
	RideHistory(int) ([]models.RideHistory, error)
}

// ServiceImpl **
type ServiceImpl struct {
	repo repository.Repository
}

// NewServiceImpl inject depedancies user repositiory
func NewServiceImpl(repo repository.Repository) Service {
	return &ServiceImpl{repo: repo}
}

// user can find nearest cab using pickup location
func (serviceImpl ServiceImpl) BookCab(req models.CabBookRequest) (resp models.CabBookResponce, err error) {
	searchCabRequest := models.SearchCab{PickLocation: req.PickUp}
	//get nearest cabID's
	foundCabs, err := serviceImpl.repo.SearchCab(searchCabRequest)
	if err != nil {
		return resp, err
	} else if len(foundCabs) == 0 {
		return resp, errors.New("no cab found in this location")
	}

	//get driver and cab details by cabID
	getCabDetails, err := serviceImpl.repo.GetCabDetails(foundCabs[0].CabId)
	if err != nil {
		return resp, err
	}

	resp.BookingID = helpers.GenrateRandomNumber()
	resp.CabNo = getCabDetails.CabNo
	resp.DriverName = getCabDetails.DriverName
	resp.DriverMob = getCabDetails.DriverMob

	return resp, nil
}

// get Ride History by UserId
func (serviceImpl ServiceImpl) RideHistory(userID int) (resp []models.RideHistory, err error) {
	resp, err = serviceImpl.repo.RideHistory(userID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
