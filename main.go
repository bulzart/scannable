package main

import (
	"artwear/controllers"
	"artwear/initializers"
	"artwear/middleware"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println(time.Now())
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Replace with your frontend's URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	r.POST("/login", controllers.Login)
	r.GET("/", controllers.Dashboard)
	r.Use(middleware.RequireAuth)
	r.POST("/signup", controllers.SignupPost)
	r.POST("create", controllers.CreateQR)

	qr := r.Group("/qr")
	qr.Use(middleware.VerifyOwner)
	qr.GET("/:id", controllers.GetQRbyId)
	qr.GET("getUserQRs", controllers.GetUserQRs)
	qr.PUT("update/:id", controllers.UpdateQR)
	qr.DELETE("delete/:id", controllers.DeleteQR)
	qr.POST("redirect/create/:id", controllers.CreateRedirect)
	qr.GET("redirect/latest/:id", controllers.GetLatestRedirect)

	err := r.Run()
	if err != nil {
		fmt.Println("Error while starting the server!")
	}
}
