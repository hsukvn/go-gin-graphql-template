package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/graphql-go/graphql"
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "alive", nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: heartbeatType,
	})

	if err != nil {
		log.Fatalf("Fail to create new schema, error: %v", err)
	}

	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema: schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":9527", nil)
}
