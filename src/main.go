package main

import (
	"github.com/gin-gonic/gin"
	"go-api/src/routes/health"
	"go-api/src/routes/api"
)

func main() {

	router := gin.Default()

	healthRouter.Init(router)
	api.Init(router)

	// start on port 8080
	router.Run(":8080")
}
