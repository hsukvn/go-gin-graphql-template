package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"firstname":  &graphql.Field{Type: graphql.String},
		"lastname":   &graphql.Field{Type: graphql.String},
		"roles":      &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// FIXME: Mockup data
				var roles []Role = []Role{
					Role{
						ID: 1,
						Name: "Research and Design",
					},
				}

				return roles, nil
			},
		},
	},
})
