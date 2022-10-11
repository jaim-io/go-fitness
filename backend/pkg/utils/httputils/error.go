package httputils

import "github.com/gin-gonic/gin"

func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Message: err.Error(),
	}
	c.JSON(status, er)
}

type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}
