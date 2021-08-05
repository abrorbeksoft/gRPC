package v1

import (
	"github.com/abrorbeksoft/gRPC/api/v1/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() *gin.Engine {
	router:=gin.Default()
	apiv1:=router.Group("v1").Use()
	apiv1.Use()
	{
		apiv1.GET("/hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, handlers.Login("Abrorbek","Ubaydullayev"))
		})
	}

	return router
}