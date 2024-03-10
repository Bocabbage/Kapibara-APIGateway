package mikanani

// TODO: code simplify

import (
	"context"
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/grpc_utils"
	"kapibara-apigateway/internal/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var some_opts []grpc.DialOption

func init() {
	// TODO: Add tls
	some_opts = append(some_opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func UpdateAnimeConfig(c *gin.Context) {
	var params UpdateAnimeFormat
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
	}

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, some_opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()

	grpcCli := grpc_utils.NewMikananiMongoCrudClient(grpcConn)

	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := grpcCli.UpdateAnime(ctx, &grpc_utils.UpdateAnimeRequest{
		Names:        []string{params.Name},
		RssUrls:      []string{params.RssUrl},
		RuleVersions: []string{params.RuleVersion},
		RuleRegexs:   []string{params.RuleVersion},
		IsActives:    []bool{true},
	})
	if err != nil {
		logger.Error()
		if status.Code(err) == codes.DeadlineExceeded {
			c.JSON(
				http.StatusGatewayTimeout,
				gin.H{"error": "The server is busy."},
			)
			return
		}

		// TODO: enrich error-handling
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Something happened at querying anime informations."},
		)
		return
	}

	if reply.SuccessCount == 1 {
		c.JSON(http.StatusOK, gin.H{"update_success": params.Name})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("unexpected update number %d", reply.SuccessCount)})
	}
}

func DelAnimeConfig(c *gin.Context) {
	delName := c.Query("del_name")
	delAllStr := c.Query("del_all")
	delAll, err := strconv.ParseBool(delAllStr)
	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "DelAnimeConfig Error: invalid param: del_all"},
		)
		return
	}

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, some_opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()

	grpcCli := grpc_utils.NewMikananiMongoCrudClient(grpcConn)

	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := grpcCli.DelAnime(ctx, &grpc_utils.DelAnimeRequest{
		DelAll: delAll,
		Names:  []string{delName},
	})
	if err != nil {
		logger.Error()
		if status.Code(err) == codes.DeadlineExceeded {
			c.JSON(
				http.StatusGatewayTimeout,
				gin.H{"error": "The server is busy."},
			)
			return
		}

		// TODO: enrich error-handling
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Something happened at querying anime informations."},
		)
		return
	}

	if reply.SuccessCount == 1 {
		c.JSON(http.StatusOK, gin.H{"del_success": delName})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("unexpected delete number %d", reply.SuccessCount)})
	}
}

func QueryAnimeConfig(c *gin.Context) {
	// TODO: Add id design for doc or add limit
	activeStr := c.Query("active_type")
	tmp, err := strconv.ParseInt(activeStr, 10, 32)
	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "QueryAnimeConfig Error: invalid param: active_type"},
		)
		return
	}
	activeType := int32(tmp)

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, some_opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()

	grpcCli := grpc_utils.NewMikananiMongoCrudClient(grpcConn)

	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := grpcCli.QueryAnime(ctx, &grpc_utils.QueryAnimeRequest{
		// TODO: current version get all
		Names:      []string{"*"},
		ActiveType: activeType,
	})
	if err != nil {
		logger.Error()
		if status.Code(err) == codes.DeadlineExceeded {
			c.JSON(
				http.StatusGatewayTimeout,
				gin.H{"error": "The server is busy."},
			)
			return
		}

		// TODO: enrich error-handling
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Something happened at querying anime informations."},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"names":  reply.Names,
			"rssUrl": reply.RssUrl,
		},
	)
}
