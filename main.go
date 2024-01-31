package main

import (
	"funderz/auth"
	"funderz/handler"
	"funderz/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect & check DB connection
	dsn := "root:@tcp(127.0.0.1:3306)/funderz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// connect repository to DB (like a model to DB in MVC)
	userRepository := user.NewRepository(db)

	// connect service to to repository (like a controller to model in MVC)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// set handler so the function in service is accessible
	userHandler := handler.NewUserHandler(userService, authService)

	// initiate API
	router := gin.Default()

	// register API
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.RegisterUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	// run router API
	router.Run()
}