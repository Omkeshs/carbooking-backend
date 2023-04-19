package repositories

import (
	"carbooking-backend/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Repository implimets all methods in Repository
type Repository interface {
	SearchCab(models.SearchCab) ([]models.FoundCab, error)
	GetCabDetails(int) (models.SearchCabResponce, error)
	RideHistory(int) ([]models.RideHistory, error)
}

// RepositoryImpl **
type RepositoryImpl struct {
	dbConn *gorm.DB
}

// NewRepositoryImpl inject dependancies of DataStore
func NewRepositoryImpl(dbConn *gorm.DB) Repository {
	return &RepositoryImpl{dbConn: dbConn}
}

// Find nearest cab list by user pickup location
func (repositoryImpl RepositoryImpl) SearchCab(searchCabReq models.SearchCab) (resp []models.FoundCab, err error) {
	dbConn := repositoryImpl.dbConn

	query := `select id as cab_id, ST_Distance_Sphere( point (?, ?), point(lng, lat)) * .001 as distance from cb_cabs having distance <= 50 order by distance asc`
	err = dbConn.Raw(query, searchCabReq.PickLocation.Longitude, searchCabReq.PickLocation.Latitude).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	if len(resp) == 0 {
		return resp, gorm.ErrRecordNotFound
	}
	return resp, nil
}

// Get Driver and Cab details by CabID
func (repositoryImpl RepositoryImpl) GetCabDetails(cab_id int) (resp models.SearchCabResponce, err error) {
	dbConn := repositoryImpl.dbConn

	err = dbConn.Table("cb_drivers").
		Select("cb_cabs.id as cab_id,cb_cabs.cab_no,cb_drivers.id as driver_id,cb_drivers.name as driver_name,cb_drivers.mobile as driver_mobile").
		Where("cb_cabs.id=?", cab_id).
		Joins("Left join cb_cabs ON cb_drivers.cab_id = cb_cabs.id").
		Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// get cab ride history
func (repositoryImpl RepositoryImpl) RideHistory(UserId int) (resp []models.RideHistory, err error) {
	dbConn := repositoryImpl.dbConn
	query := `select ride_id,user_id,cab_no,driver_id,booking_id,pickup_loc,drop_loc,created_at from cb_rides where user_id=?;`
	err = dbConn.Raw(query, UserId).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	if len(resp) == 0 {
		return resp, gorm.ErrRecordNotFound
	}
	return resp, nil
}
