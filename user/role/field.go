package role

import (
	"github.com/graphql-go/graphql"
)

func RolesQueryFieldGet() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(roleType),
		Resolve: rolesResolveGet,
	}
}
