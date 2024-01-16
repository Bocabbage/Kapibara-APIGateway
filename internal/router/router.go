package router

import (
	"kapibara-apigateway/internal/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ServerEngine *gin.Engine

func init() {
	ServerEngine = gin.Default()
	ServerEngine.Use(cors.New(cors.Config{
		// [todo] restrict the cors
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	authRouter := ServerEngine.Group("/auth")
	{
		authRouter.POST("/login", auth.AuthLogin)
	}
}
