package router

import (
	"fmt"
	"kapibara-apigateway/internal/auth"
	"kapibara-apigateway/internal/logger"

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
		authRouter.POST("/register", auth.AuthRegister)
	}

	logger.Debug("Gin router init finished.")
	fmt.Println("Gin router init finished.")
}
