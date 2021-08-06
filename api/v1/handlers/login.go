package handlers

import (
	"github.com/abrorbeksoft/gRPC/api/v1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (user *handlerV1) Login(c *gin.Context)  {
	c.JSON(http.StatusOK,models.User{
		Name: "Abrorbek",
		Surname: "Ubaydullayev",
		Age: 21,
	})
}
