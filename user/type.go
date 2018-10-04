package user

import (
	"github.com/graphql-go/graphql"
	"github.com/hsukvn/go-graphql-template/user/role"
)

type user struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

// FIXME: Mockup data
var users []user = []user{
	user{
		ID: 1,
		Firstname: "Kevin",
		Lastname: "Hsu",
	},
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"firstname":  &graphql.Field{Type: graphql.String},
		"lastname":   &graphql.Field{Type: graphql.String},
		"roles":      role.RolesQueryFieldGet(),
	},
})
