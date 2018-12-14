package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(c *Config) (*Server, error) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		router: newRouter(&RouterConfig{
			DisableAuth: c.DisableAuth,
		}),
	}, nil
}

func (s *Server) Run(port int) {
	p := strconv.Itoa(port)
	s.router.Run(":" + p)
}
