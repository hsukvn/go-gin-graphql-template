package role

import (
	"github.com/graphql-go/graphql"
)

type role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FIXME: Mockup data
var roles []role = []role{
	role{
		ID: 1,
		Name: "Research and Design",
	},
}

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})
