package models

import "time"

type Location struct {
	Area      string  `json:"Area" binding:"max=20"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type CabBookRequest struct {
	UserId int32 `json:"user_id" binding:"gte=1"`
	Time   string
	PickUp Location `json:"pickup_loc"`
	Drop   Location `json:"drop_loc"`
}

type CabBookResponce struct {
	BookingID  int16  `json:"booking_id"`
	DriverName string `json:"driver_name"`
	DriverMob  string `gorm:"column:driver_mobile" json:"driver_mob"`
	CabNo      string `json:"cab_no"`
}

type RideHistoryRequest struct {
	UserID int `json:"user_id" binding:"required,gte=1"`
}

type RideHistory struct {
	RideId    int       `json:"ride_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"user_id"`
	CabNo     string    `json:"cab_no" gorm:"cab_no"`
	DriverId  int       `json:"driver_id" gorm:"driver_id"`
	BookingID int       `json:"booking_id" gorm:"booking_id"`
	PickUpLoc string    `gorm:"column:pickup_loc" json:"pickup_loc" `
	DropLoc   string    `gorm:"drop_loc" json:"drop_loc" `
	DateTime  time.Time `gorm:"DateTime" json:"DateTime"`
}

type SearchCab struct {
	PickLocation Location
}

type FoundCab struct {
	CabId    int     `gorm:"cab_id"`
	Distance float64 `gorm:"distance"`
}

type SearchCabResponce struct {
	CabID      int32  `gorm:"cab_id"`
	CabNo      string `gorm:"cab_no"`
	DriverId   int32  `gorm:"driver_id"`
	DriverName string `gorm:"driver_name"`
	DriverMob  string `gorm:"column:driver_mobile"`
}
