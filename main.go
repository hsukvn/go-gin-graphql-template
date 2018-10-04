package main

import (
	"net/http"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hsukvn/go-graphql-template/heartbeat"
	"github.com/hsukvn/go-graphql-template/user"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"heartbeat": heartbeat.HeartbeatQueryFieldGet(),
				"users":     user.UsersQueryFieldGet(),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: graphql.Fields{
				"createUser": user.CreateUserMutationFieldGet(),
			},
		}),
	})

	if err != nil {
		log.Fatalf("Fail to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/", h)
	http.ListenAndServe(":9527", nil)
}
