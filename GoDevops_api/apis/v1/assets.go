package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAssets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "all assets",
		"data": gin.H{},
	})
}
