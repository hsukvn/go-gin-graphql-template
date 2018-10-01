package main

import (
	"net/http"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Heartbeat struct {
	Status string `json:"status"`
}

func main() {
	heartbeatType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Heartbeat",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"heartbeat": &graphql.Field{
				Type: heartbeatType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Heartbeat{
						Status: "alive",
					}, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
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
