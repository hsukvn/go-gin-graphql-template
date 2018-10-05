package controller

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hsukvn/go-gin-graphql-template/graphql/user"
)

type GraphqlController struct{}

func (ctr *GraphqlController) newSchema() (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": user.GetUsersQueryField(),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: graphql.Fields{
				"createUser": user.GetCreateUserMutationField(),
			},
		}),
	})

	return schema, err
}

func (ctr *GraphqlController) NewHandler() (*handler.Handler) {
	schema, err := ctr.newSchema()

	if err != nil {
		log.Fatalf("Fail to create schema, error: %v", err)
		return nil
	}

	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
}
