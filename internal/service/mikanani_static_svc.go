package service

import (
	"fmt"
	"kapibara-apigateway/internal/config"
	"kapibara-apigateway/internal/logger"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var opts []grpc.DialOption

func init() {
	// TODO: Add tls
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		logger.Error(fmt.Sprintf("[SaveUploadedFile][FAILED]: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	logger.Info(fmt.Sprintf("[SaveUploadedFile][SUCCESS]: uid[%s]", uidStr))
	c.JSON(http.StatusOK, gin.H{})
}
