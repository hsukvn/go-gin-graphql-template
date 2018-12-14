package server

import (
	"fmt"
	"log"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/hsukvn/go-gin-graphql-template/controller"
	"github.com/hsukvn/go-gin-graphql-template/lib/user"
)

const identityKey = "name"
const secretKey = "somesecretkey"

type login struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required"`
}

type User struct {
	Name string
}

func getAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Name: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			name := loginVals.Name
			passwd := loginVals.Passwd

			if err := user.AuthenticateUser(name, passwd); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &User{
				Name: name,
			}, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
		return nil, fmt.Errorf("auth: Fail to setup jwt auth middleware, err: (%v)", err)
	}

	return authMiddleware, nil
}

func newRouter(c *RouterConfig) *gin.Engine {
	r := gin.Default()

	authMiddleware, err := getAuthMiddleware()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	ping := new(controller.PingController)

	r.GET("/ping", ping.GetController)
	r.POST("/login", authMiddleware.LoginHandler)

	r.Use(authMiddleware.MiddlewareFunc())

	graphql := new(controller.GraphqlController)

	r.POST("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	return r
}
