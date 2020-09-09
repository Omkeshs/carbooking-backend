package models

type Location struct {
	Latitude  float64
	Longitude float64
}

type CabBookRequest struct {
	UserID int
	Time   string
	PickUp Location
	Drop   Location
}

type CabBookResponce struct {
	BookingID  int
	DriverName string
	DriverMob  string
	CabNo      string
}

type RideHistoryRequest struct {
	UserID int
}

type RideHistory struct {
	UserID    int
	DateTime  string
	Locations []Location
	CabNo     string
}
