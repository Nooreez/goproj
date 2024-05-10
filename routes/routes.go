package routes

import (
	"goproj/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	//	Authorization
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/logout", controllers.Logout)
	r.POST("/reset-password", controllers.ResetPassword)
	//	Menu
	r.GET("/showmenu", controllers.ShowMenu)
	r.POST("/addmenu", controllers.AddMenu)
	r.DELETE("/deletemenu", controllers.DeleteMenu)
}
