package main

import (
	"net/http"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hsukvn/go-graphql-template/queries"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: queries.GetRootFields(),
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
