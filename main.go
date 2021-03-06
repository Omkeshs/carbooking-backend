package main

import (
	"fmt"
	"time"

	endpoints "github.com/Omkeshs/carbooking-backend/endpoints"

	"github.com/Omkeshs/carbooking-backend/database"

	service "github.com/Omkeshs/carbooking-backend/service"

	repository "github.com/Omkeshs/carbooking-backend/repositories"

	handler "github.com/Omkeshs/carbooking-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	HTTPAddr = "8080"
)

func main() {
	router := gin.Default()

	//CORS ISSUE
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	conPool := database.NewMysqlConnection()
	taskRepository := repository.NewRepositoryImpl(conPool)
	taskService := service.NewServiceImpl(taskRepository)
	taskHandler := handler.NewHandlerImpl(taskService)
	endpoints.NewRoute(router, taskHandler)

	fmt.Println(router.Run(":" + HTTPAddr))
}
