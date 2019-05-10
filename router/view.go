package router

import (
	"API1/sd"
	"github.com/gin-gonic/gin"
)



func Load(g *gin.Engine) *gin.Engine{
	svcd := g.Group("/v1")
	{
		svcd.GET("/health",sd.HealthCheck)
		svcd.GET("/disk",sd.DiskCheck)
		svcd.GET("/cpu",sd.CPUCheck)
		svcd.GET("/ram",sd.RAMCheck)
	}
	return  g
}