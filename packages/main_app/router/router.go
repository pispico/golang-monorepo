package router

import (
	"github.com/gin-gonic/gin"
	health "github.com/pispico/golang-monorepo/packages/shared/handlers/health"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Main app"))
	return r
}
