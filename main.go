package main

import (
	"fmt"
	"log"
	"ugbisa/auth"
	"ugbisa/handler"
	"ugbisa/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:cimanggis123@tcp(127.0.0.1:3306)/pplug?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.eJxZ8r3eJvpSYy4Y4jjw1g0fjO6jcgTDOx5XJ7v8K8c")
	if err != nil {
		fmt.Println("Error")
	}

	if token.Valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// input
	// handler mapping input ke struct
	// service mapping ke struct User
	// repository save struct User ke db
	// db
}
