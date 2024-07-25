package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"life-hub-backend/config"
)

func Handle() {
	r := gin.Default()

	api := r.Group("/api")
	authentication := api.Group("/auth")

	HandleAuthentication(authentication)

	err := r.Run(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		panic("Couldn't start server")
	}
}
