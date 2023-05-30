package http

import "github.com/gin-gonic/gin"

func router(releaseMode, useLogger, useRecovery bool) *gin.Engine {
	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	if useLogger {
		g.Use(gin.Logger())
	}
	if useRecovery {
		g.Use(gin.Recovery())
	}
	return g
}
