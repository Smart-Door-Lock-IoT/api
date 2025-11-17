package http

import (
	_ "github.com/Smart-Door-Lock-IoT/api/docs"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration"
	"github.com/Smart-Door-Lock-IoT/api/features/control"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer() {
	router := gin.Default()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/api/v1")

	configuration.RegisterConfiguration(v1)
	control.RegisterControl(v1)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	_ = router.Run(":8080")
}
