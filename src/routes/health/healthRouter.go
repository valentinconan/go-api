package healthRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthRouter struct {
}

func NewHealthRouter() *HealthRouter {
	return &HealthRouter{}
}

func (healthRouter *HealthRouter) Init(router *gin.RouterGroup) {

	router.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state":   "UP",
			"version": "master"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.Status(200)
	})
}
