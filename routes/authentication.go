package routes

import (
	"github.com/gin-gonic/gin"
	"life-hub-backend/controllers"
)

func HandleAuthentication(authentication *gin.RouterGroup) {
	authentication.POST("/sign-up", controllers.SignUp)
	authentication.POST("/sign-up-google", controllers.SignUpWithGoogle)
	authentication.POST("/sign-in", controllers.SignIn)
}
