package main

import (
	"artwear/controllers"
	"artwear/initializers"
	"artwear/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	r.Use(middleware.RequireAuth)
	r.POST("/signup", controllers.SignupPost)
	r.POST("create", controllers.CreateQR)
	qr := r.Group("/qr")
	qr.Use(middleware.VerifyOwner)
	qr.GET("getUserQRs", controllers.GetUserQRs)
	qr.PUT("update/:id", controllers.UpdateQR)
	qr.DELETE("delete/:id", controllers.DeleteQR)
	err := r.Run()
	if err != nil {
		fmt.Println("Error while starting the server!")
	}
}