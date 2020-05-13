package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vincecfl/dex-robot/handler"
	"github.com/vincecfl/dex-robot/router/middleware"
	"net/http"
)

const (
	IsAdmin = int(1)
	IsUser  = int(2)
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The Incorrect API Route")
	})

	check := g.Group("dex/api/check")
	{
		check.GET("/health", handler.HealthCheck)
		check.GET("/disk", handler.DiskCheck)
		check.GET("/cpu", handler.CPUCheck)
		check.GET("/ram", handler.RAMCheck)
	}

	return g
}
