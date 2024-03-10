package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestEcho(c *gin.Context) {
	msg := c.Query("message")
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
