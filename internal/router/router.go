package router

import (
	"kapibara-apigateway/internal/auth"
	"kapibara-apigateway/internal/logger"
	"kapibara-apigateway/internal/middlewares"
	"kapibara-apigateway/internal/mikanani"
	"kapibara-apigateway/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ServerEngine *gin.Engine

func init() {
	ServerEngine = gin.Default()
	ServerEngine.Use(cors.New(cors.Config{
		// [todo] restrict the cors
		// AllowAllOrigins:  true,
		AllowOrigins:     []string{"http://192.168.4.29", "http://kapibara.local.com:5173"}, // [todo] dev-mode
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST"},
	}))

	authRouter := ServerEngine.Group("/auth/v1")
	{
		authRouter.POST("/login", auth.AuthLogin)
		authRouter.POST("/register", auth.AuthRegister)
	}

	mikananiCrudRouter := ServerEngine.Group("/mikanani/v1")
	mikananiCrudRouter.Use(middlewares.TokenValidationMid())
	{
		mikananiCrudRouter.GET("/query", mikanani.QueryAnimeConfig)
		mikananiCrudRouter.PUT("/update", mikanani.UpdateAnimeConfig)
		mikananiCrudRouter.PUT("/delete", mikanani.DelAnimeConfig)
	}

	testRouter := ServerEngine.Group("/test/v1")
	testRouter.Use(middlewares.TokenValidationMid())
	{
		testRouter.GET("/echo", utils.TestEcho)
	}

	logger.Debug("Gin router init finished.")
}
