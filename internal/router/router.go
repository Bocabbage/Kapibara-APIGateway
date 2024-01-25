package router

import (
	"kapibara-apigateway/internal/auth"
	"kapibara-apigateway/internal/logger"
	"kapibara-apigateway/internal/middlewares"
	"kapibara-apigateway/internal/utils"

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

	authRouter := ServerEngine.Group("/auth/v1")
	{
		authRouter.POST("/login", auth.AuthLogin)
		authRouter.POST("/register", auth.AuthRegister)
	}

	testRouter := ServerEngine.Group("/test/v1")
	testRouter.Use(middlewares.TokenValidationMid())
	{
		testRouter.GET("/echo", utils.TestEcho)
	}

	logger.Debug("Gin router init finished.")
}
