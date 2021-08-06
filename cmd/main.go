package main

import (
	v1 "github.com/abrorbeksoft/gRPC/api"
	"github.com/gin-gonic/gin"
)

func main() {
	var routers *gin.Engine

	routers= v1.Main()

	routers.Run(":3000")
}
