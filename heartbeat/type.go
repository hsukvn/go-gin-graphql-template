package heartbeat

import (
	"github.com/graphql-go/graphql"
)

type heartbeat struct {
	Status string `json:"status"`
}

var heartbeatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Heartbeat",
	Fields: graphql.Fields{
		"status": &graphql.Field{Type: graphql.String},
	},
})
