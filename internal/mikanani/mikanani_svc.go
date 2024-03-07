package mikanani

import (
	"context"
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	mikansvc "kapibara-apigateway/internal/mikanani_grpc_utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var opts []grpc.DialOption

func init() {
	// TODO: Add tls
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// @Summary List meta data of animes
// @Description Get anime meta data in a range
// @Accept json
// @Produce json
// @Param start int true 			"startIndex"
// @Param end int true 				"endIndex"
// @Params status_filter int true 	"statusFilter"
// @Success 200 {array} AnimeMeta
// @Failure 400 {object} ClientErrorResponse
// @Failure 500 {object} ServerErrorResponse
// @Router /mikanani/v2/anime/list-meta [get]
func ListAnimeMeta(c *gin.Context) {
	var params ListAnimeMetaFormat
	if c.ShouldBind(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()
	cli := mikansvc.NewMikananiServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := cli.ListAnimeMeta(ctx, &mikansvc.ListAnimeMetaRequest{
		StartIndex:   params.StartIdx,
		EndIndex:     params.EndIdx,
		StatusFilter: params.StatusFilter,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[ListAnimeMeta][Error]: %v", err))
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
			"count": reply.GetItemCount(),
			"metas": reply.GetAnimeMetas(),
		},
	)

}

// @Summary Get anime doc
// @Description Get doc of an anime
// @Accept json
// @Produce json
// @Param uid int true 				Anime-uid
// @Success 200 {object} AnimeDoc
// @Failure 400 {object} ClientErrorResponse
// @Failure 500 {object} ServerErrorResponse
// @Router /mikanani/v2/anime/doc [get]
func GetAnimeDoc(c *gin.Context) {
	var params GetAnimeDocFormat
	if c.ShouldBind(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()
	cli := mikansvc.NewMikananiServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := cli.GetAnimeDoc(ctx, &mikansvc.GetAnimeDocRequest{
		Uid: params.Uid,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[GetAnimeDoc][Error]: %v", err))
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
			"doc": reply.GetAnimeDoc(),
		},
	)
}

func UpdateAnimeDoc(c *gin.Context) {
	var params UpdateAnimeDocFormat
	if c.ShouldBindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	grpcConn, err := grpc.Dial(config.GlobalConfig.MikananiConf.GRpcServerAddr, opts...)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "backend service access failed."},
		)
		return
	}
	defer grpcConn.Close()
	cli := mikansvc.NewMikananiServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), config.GRPC_CLI_TIMEOUT)
	defer cancel()

	reply, err := cli.UpdateAnimeDoc(ctx, &mikansvc.UpdateAnimeDocRequest{
		UpdateAnimeDoc: &mikansvc.AnimeDoc{
			Uid:    params.Uid,
			RssUrl: params.RssUrl,
			Rule:   params.Rule,
			Regex:  params.Regex,
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[UpdateAnimeDoc][Error]: %v", err))
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
		gin.H{},
	)
}

func UpdateAnimeMeta(c *gin.Context) {
	var params UpdateAnimeMetaFormat
	if c.ShouldBindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
}

func InsertAnimeItem(c *gin.Context) {
	var params InsertAnimeItemFormat
	if c.ShouldBindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
}

func DeleteAnimeItem(c *gin.Context) {
	var params DeleteAnimeItemFormat
	if c.ShouldBind(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
}

func DispatchDownload(c *gin.Context) {

}
