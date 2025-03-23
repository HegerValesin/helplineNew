package main

import (
	"github.com/gin-gonic/gin"
	"helplineBKgo/database"
	"helplineBKgo/handlers"
	"helplineBKgo/repositories"
	"helplineBKgo/services"
)

func main() {
	database.InitDB()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.List)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
