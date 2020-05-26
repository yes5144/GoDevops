package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// fake data

// versions :=[

// 	["nz","wx","1.0.21,32"],
// 	["nz","hf","1.0.21.132"]

// ]

// GetAllVersions xxx
func GetAllVersions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "all versions,",
		"data": gin.H{},
	})
}
