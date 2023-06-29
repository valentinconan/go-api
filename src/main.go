package main

import (
	"github.com/gin-gonic/gin"
	"go-api/src/routes/health"
)

func main() {

	router := gin.Default()

	healthRouter.Init(router)

	// start on port 8080
	router.Run(":8080")
}
