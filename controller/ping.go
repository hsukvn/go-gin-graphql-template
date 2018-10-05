package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (ctr *PingController) GetController(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
