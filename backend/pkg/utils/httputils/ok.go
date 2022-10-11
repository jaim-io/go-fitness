package httputils

import "github.com/gin-gonic/gin"

func NewOk(c *gin.Context, status int) {
	ok := HTTPOK{
		Message: "status OK",
	}
	c.JSON(status, ok)
}

type HTTPOK struct {
	Message string `json:"message" example:"status OK"`
}
