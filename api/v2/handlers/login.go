package handlers

import (
	"github.com/abrorbeksoft/gRPC/api/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handlerV2 *handlerV2) Login(c *gin.Context)  {
	c.JSON(http.StatusOK,models.User{
		Id: c.Query("id"),
		Name: c.Query("name"),
		Surname: c.Query("surname"),
		Age: c.Query("age"),
		Salary: c.Query("salary"),
	})
}


