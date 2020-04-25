package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yes5144/GoDevops/GoDevops_api/apis/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mes":  "pong",
			"data": "no data",
		})
	})

	apiv1 := r.Group("/api/v1")
	{
		// assets
		assets := apiv1.Group("/assets")
		{
			assets.GET("/", v1.GetAllAssets)
		}
	}

	return r
}
