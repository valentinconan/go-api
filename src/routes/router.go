package router

import (
	"github.com/gin-gonic/gin"
	"go-api/src/routes/api"
	"go-api/src/routes/health"
	"log/slog"
)

func Init() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	// catch all panic exception and return 500
	router.Use(gin.Recovery())

	//base endpoint not secured
	rootGroup := router.Group("/")

	healthRouter.NewHealthRouter().Init(rootGroup)

	//secured group
	apiGroup := router.Group("/api")
	apiGroup.Use(checkAuthorization())
	api.NewApiRouter().Init(apiGroup)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Erreur lors du lancement du serveur Gin: %v", err)
	}
}
func checkAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Error("All security check must be proceed here")
	}
}
