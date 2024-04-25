package router

import (
	"fmt"
	"kapibara-apigateway/internal/auth"
	"kapibara-apigateway/internal/config"
	kgrpc "kapibara-apigateway/internal/grpc"
	"kapibara-apigateway/internal/logger"
	"kapibara-apigateway/internal/middlewares"
	"kapibara-apigateway/internal/service"
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
	ServerEngine.MaxMultipartMemory = 8 << 20 // 8 MiB

	authRouter := ServerEngine.Group("/auth/v1")
	{
		authRouter.POST("/login", auth.AuthLogin)
		authRouter.POST("/register", auth.AuthRegister)
	}

	mikananiServiceRouter := ServerEngine.Group("/mikanani/v2")
	mikananiServiceRouter.Use(middlewares.TokenValidationMid())
	{
		// Static sources APIs
		mikananiServiceRouter.GET("/pics/:uid", service.GetAnimeImage)
		mikananiServiceRouter.POST("/pics/upload/:uid", service.PostAnimeImage)

		animeServiceRouter := mikananiServiceRouter.Group("/anime")
		{
			// Http -> gRPC
			mikananiMux, err := kgrpc.CreateMikananiMux(config.GlobalConfig.MikananiConf.GRpcServerAddr)
			if err != nil {
				logger.Fatal(fmt.Sprintf("[CreateMikananiMux]Error: %v", err))
			}
			animeServiceRouter.Any("/*any", gin.WrapH(mikananiMux))
		}
	}

	testRouter := ServerEngine.Group("/test/v1")
	testRouter.Use(middlewares.TokenValidationMid())
	{
		testRouter.GET("/echo", utils.TestEcho)
	}

	logger.Debug("Gin router init finished.")
}
