package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yes5144/GoDevops/GoDevops_api/apis/v1"
	"github.com/yes5144/GoDevops/GoDevops_api/controllers"
	"github.com/yes5144/GoDevops/GoDevops_api/middlewares"
)

// InitRouter xxx
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mes":  "pong",
			"data": "no data",
		})
	})

	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	r.GET("/api/auth/info", middlewares.JwtMiddleware(), controllers.Info)

	apiv1 := r.Group("/api/v1", middlewares.JwtMiddleware())
	{
		// assets
		assets := apiv1.Group("/assets")
		{
			assets.GET("/", v1.GetAllAssets)
		}

		// versions
		versions := apiv1.Group("/versions")
		{
			versions.GET("/", v1.GetAllVersions)
		}
	}

	return r
}
