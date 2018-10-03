package types

import (
	"github.com/graphql-go/graphql"
)

type Heartbeat struct {
	Status string `json:"status"`
}

var HeartbeatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Heartbeat",
	Fields: graphql.Fields{
		"status": &graphql.Field{Type: graphql.String},
	},
})
