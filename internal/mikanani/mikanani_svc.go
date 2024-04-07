package mikanani

import (
	"context"
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	mikansvc "kapibara-apigateway/internal/mikanani_grpc_utils"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	if c.ShouldBindQuery(&params) != nil {
		logger.Debug(fmt.Sprintf("query params: %v\n", params))
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
		logger.Debug(fmt.Sprintf("query grpc error, params: %v\n", params))
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
			gin.H{"error": "Something happened at querying anime doc informations."},
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

	intUid, err := strconv.ParseInt(params.Uid, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid uid."},
		)
		return
	}
	_, err = cli.UpdateAnimeDoc(ctx, &mikansvc.UpdateAnimeDocRequest{
		UpdateAnimeDoc: &mikansvc.AnimeDoc{
			Uid:    intUid,
			RssUrl: params.RssUrl,
			Rule:   params.Rule,
			Regex:  params.Regex,
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[UpdateAnimeDoc][Error]: %v", err))
		logger.Debug(fmt.Sprintf("[UpdateAnimeDoc][Error]params: %v", params))
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
			gin.H{"error": "Something happened at updating anime doc informations."},
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

	intUid, err := strconv.ParseInt(params.Uid, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid uid."},
		)
		return
	}

	_, err = cli.UpdateAnimeMeta(ctx, &mikansvc.UpdateAnimeMetaRequest{
		UpdateAnimeMeta: &mikansvc.AnimeMeta{
			Uid:            intUid,
			Name:           params.Name,
			DownloadBitmap: params.DownloadBitmap,
			IsActive:       params.IsActive,
			Tags:           params.Tags,
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[UpdateAnimeMeta][Error]: %v", err))
		logger.Debug(fmt.Sprintf("[UpdateAnimeMeta][Error]params: %v", params))
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
			gin.H{"error": "Something happened at updating anime meta informations."},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

func InsertAnimeItem(c *gin.Context) {
	var params InsertAnimeItemFormat
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

	res, err := cli.InsertAnimeItem(ctx, &mikansvc.InsertAnimeItemRequest{
		InsertAnimeMeta: &mikansvc.AnimeMeta{
			Uid:            -1,
			Name:           params.Name,
			DownloadBitmap: 0,
			IsActive:       params.IsActive,
			Tags:           params.Tags,
		},
		InsertAnimeDoc: &mikansvc.AnimeDoc{
			Uid:    -1,
			RssUrl: params.RssUrl,
			Rule:   params.Rule,
			Regex:  params.Regex,
		},
	})

	if err != nil {
		logger.Error(fmt.Sprintf("[InsertAnimeItem][Error]: %v", err))
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
			gin.H{"error": "Something happened at inserting anime informations."},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"uid": res.Uid},
	)

}

func DeleteAnimeItem(c *gin.Context) {
	var params DeleteAnimeItemFormat
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

	_, err = cli.DeleteAnimeItem(ctx, &mikansvc.DeleteAnimeItemRequest{Uid: params.Uid})

	if err != nil {
		logger.Error(fmt.Sprintf("[DeleteAnimeItem][Error]: %v", err))
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
			gin.H{"error": "Something happened at deleting anime informations."},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

func DispatchDownload(c *gin.Context) {
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

	_, err = cli.DispatchDownloadTask(ctx, &emptypb.Empty{})
	if err != nil {
		logger.Error(fmt.Sprintf("[DispatchDownload][Error]: %v", err))
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
			gin.H{"error": "Something happened at dispatching download task."},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{},
	)
}

func GetAnimeImage(c *gin.Context) {
	uidStr := c.Param("uid")
	uidStr = strings.Trim(uidStr, "/")
	rpath := config.GlobalConfig.MountConf.MikananiNFSMountPath
	imagePath := filepath.Join(rpath, "/pics", fmt.Sprintf("%s.png", uidStr))
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Invalid uid: %s", uidStr)},
		)
		return
	}

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.String(http.StatusOK, string(imageData))
}

func PostAnimeImage(c *gin.Context) {
	uidStr := c.Param("uid")
	uidStr = strings.Trim(uidStr, "/")
	rpath := config.GlobalConfig.MountConf.MikananiNFSMountPath
	imagePath := filepath.Join(rpath, "/pics", fmt.Sprintf("%s.png", uidStr))

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only accept png."})
		return
	}

	c.SaveUploadedFile(file, imagePath)
	c.JSON(http.StatusOK, gin.H{})
}
