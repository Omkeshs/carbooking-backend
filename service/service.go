package services

import (
	"fmt"
	"std/omkesh/carBooking-Backend/models"
	repository "std/omkesh/carBooking-Backend/repositories"
)

// Service describes the service.
type Service interface {
	BookCab(models.CabBookRequest) (*models.CabBookResponce, error)
}

//ServiceImpl **
type ServiceImpl struct {
	repo repository.Repository
}

//NewServiceImpl inject depedancies user repositiory
func NewServiceImpl(repo repository.Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (serviceImpl ServiceImpl) BookCab(req models.CabBookRequest) (resp *models.CabBookResponce, err error) {
	fmt.Println("STEP 3 : FROM SERVICE")
	resp, _ = serviceImpl.repo.BookCab(req)
	return resp, nil
}
