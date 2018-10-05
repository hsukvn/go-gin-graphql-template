package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (pc PingController) GetController(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
