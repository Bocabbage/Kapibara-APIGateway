package router

import (
	"kapibara-apigateway/internal/auth"
	"kapibara-apigateway/internal/config"
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
	ServerEngine.SetTrustedProxies(nil)
	ServerEngine.Use(cors.New(cors.Config{
		// [todo] restrict the cors
		// AllowAllOrigins:  true,
		AllowOrigins:     config.GlobalConfig.CORSConf.AllowOrigins, // [todo] dev-mode
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
	}))

	authRouter := ServerEngine.Group("/auth/v1")
	{
		authRouter.POST("/login", auth.AuthLogin)
		authRouter.POST("/register", auth.AuthRegister)
	}

	mikananiServiceRouter := ServerEngine.Group("/mikanani/v2")
	mikananiServiceRouter.Use(middlewares.TokenValidationMid())
	{
		animeServiceRouter := mikananiServiceRouter.Group("/anime")
		{
			animeServiceRouter.GET("/list-meta", mikanani.ListAnimeMeta)
			animeServiceRouter.GET("/doc", mikanani.GetAnimeDoc)
			animeServiceRouter.PUT("/update-doc", mikanani.UpdateAnimeDoc)
			animeServiceRouter.PUT("/update-meta", mikanani.UpdateAnimeMeta)
			animeServiceRouter.POST("/insert", mikanani.InsertAnimeItem)
			animeServiceRouter.DELETE("/delete", mikanani.DeleteAnimeItem)
			animeServiceRouter.POST("/dispatch-download", mikanani.DispatchDownload)
		}
	}

	testRouter := ServerEngine.Group("/test/v1")
	testRouter.Use(middlewares.TokenValidationMid())
	{
		testRouter.GET("/echo", utils.TestEcho)
	}

	logger.Debug("Gin router init finished.")
}
