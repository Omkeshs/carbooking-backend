package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Create User if not exist
// CREATE USER 'car_user'@'localhost' IDENTIFIED BY 'password';
// GRANT ALL ON *.* TO 'car_user'@'localhost';
// sudo mysql -ucar_user -ppassword car_book < CarBookingPortal.sql

const (
	MYSQLDBUSER = "car_user"
	MYSQLDBPASS = "password"
	MYSQLDBADDR = "car_book"
)

// NewMysqlConnection for client connection
func NewMysqlConnection() *gorm.DB {

	connectionStr := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		MYSQLDBUSER,
		MYSQLDBPASS,
		MYSQLDBADDR,
	)

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", connectionStr)
	if err != nil {
		panic("Error In Create Client Connection" + err.Error())
	}
	return client
}
