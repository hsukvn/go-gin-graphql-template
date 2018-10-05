package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hsukvn/go-gin-graphql-template/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)
	graphql := new(controller.GraphqlController)

	r.GET("/ping", ping.GetController)
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	return r
}
