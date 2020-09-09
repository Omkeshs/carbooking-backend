package main

import (
	"fmt"
	"std/omkesh/carBooking-Backend/database"
	"std/omkesh/carBooking-Backend/endpoints"

	handler "std/omkesh/carBooking-Backend/handlers"
	repository "std/omkesh/carBooking-Backend/repositories"
	service "std/omkesh/carBooking-Backend/service"

	"github.com/gin-gonic/gin"
)

const (
	HTTPAddr = "8080"
)

func main() {
	fmt.Println("STEP 1 : FROM MAIN")
	router := gin.Default()

	conPool := database.NewMysqlConnection()
	taskRepository := repository.NewRepositoryImpl(conPool)
	taskService := service.NewServiceImpl(taskRepository)
	taskHandler := handler.NewHandlerImpl(taskService)
	endpoints.NewRoute(router, taskHandler)

	fmt.Println(router.Run(":" + HTTPAddr))
}
