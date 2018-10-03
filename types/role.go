package types

import (
	"github.com/graphql-go/graphql"
)

type Role struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})
