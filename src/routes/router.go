package router

import (
    "log"
	"github.com/gin-gonic/gin"
	"go-api/src/routes/health"
	"go-api/src/routes/api"
)

func Init(){

    router := gin.Default()
    gin.SetMode(gin.ReleaseMode)
    // catch all panic exception and return 500
    router.Use(gin.Recovery())

    //base endpoint not secured
    rootGroup := router.Group("/")

    healthRouter.Init(rootGroup)

    //secured group
    apiGroup := router.Group("/api")
    apiGroup.Use(checkAuthorization())
    api.Init(apiGroup)

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Erreur lors du lancement du serveur Gin: %v", err)
    }
}
func checkAuthorization() gin.HandlerFunc {
   return func(c *gin.Context) {
      log.Println("All security check must be proceed here")
   }
}