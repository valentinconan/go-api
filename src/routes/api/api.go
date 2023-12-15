package api

import (
	"github.com/gin-gonic/gin"
	"go-api/src/routes/tools"
	"log/slog"
	"net/http"
	"strconv"
)

type Router struct {
	Service tools.IHttpService
}

func NewApiRouter() *Router {
	return &Router{
		tools.NewHttpService(),
	}
}

func (r *Router) Init(router *gin.RouterGroup) {

	router.GET("/:id", validatePathParamFormat(), func(c *gin.Context) {

		id := c.Param("id")

		c.Header("Content-type", "application/json")
		c.Header("charset", "UTF-8")
		c.JSON(http.StatusOK, gin.H{
			"level":   "INFO",
			"message": "Path param is correct",
			"value":   "The value of the path param is: " + id})
	})

	router.GET("/retrieve/:action", func(c *gin.Context) {

		action := c.Param("action")

		resp := r.Service.Call("http://localhost:8080/" + action)

		c.Header("Content-type", "application/json")
		c.Header("charset", "UTF-8")

		c.JSON(http.StatusOK, gin.H{
			"level":   "INFO",
			"message": "Path param is correct",
			"value":   "The value of the request is: " + resp})
	})
}

// Middleware in order to validate path param
func validatePathParamFormat() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		if isInt(id) != true {
			slog.Error("Error: The parameter is not an integer")
			c.AbortWithStatus(http.StatusBadRequest)
			c.Header("Content-type", "application/json")
			c.JSON(http.StatusOK, gin.H{
				"level":   "ERROR",
				"message": "path param incorrect. It must be an integer",
				"value":   "The value of the path param is: " + id})
			return
		}

		c.Next()
	}
}
func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
