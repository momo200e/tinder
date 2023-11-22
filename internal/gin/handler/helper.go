package handler

import (
	"fmt"
	"net/http"
	"time"

	"tinder/domain"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func Failed(c *gin.Context, err domain.ErrorFormat, customMessage string) {
	message := err.Message
	if customMessage != "" {
		message = customMessage
	}

	switch err {
	case domain.ErrorServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	case domain.ErrorBadRequest:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    err.Code,
			"message": message,
			"time":    fmt.Sprintf("%d", time.Now().Unix()),
		})
	}
}
