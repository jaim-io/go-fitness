package controllers

import (
	_ "github.com/Jaim010/jaim-io/backend/pkg/utils/httputils"
	"github.com/gin-gonic/gin"
)

// GetHealth godoc
// @Summary     Get server health
// @Description get server health
// @Tags        health
// @Accept      json
// @Produce     json
// @Success     200 {object}   	httputil.HTTPOK
// @Router      /health [get]
func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "healthy",
	})
}
