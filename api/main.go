package api

import (
	handlerV1 "github.com/abrorbeksoft/gRPC/api/v1/handlers"
	handlerV2 "github.com/abrorbeksoft/gRPC/api/v2/handlers"

	"github.com/gin-gonic/gin"
)

func Main() *gin.Engine {
	router:=gin.Default()
	apiv1:=router.Group("v1").Use()
	handlerV1:=handlerV1.New()

	// apiv1 methodlar
	apiv1.Use()
	{
		apiv1.GET("/hello", handlerV1.Login)
	}
	apiv2:=router.Group("/v2").Use()
	handlerV2:=handlerV2.New()

	apiv2.Use()
	{
		apiv2.GET("/login",handlerV2.Login)
	}
	return router
}