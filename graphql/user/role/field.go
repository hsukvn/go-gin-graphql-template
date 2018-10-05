package role

import (
	"github.com/graphql-go/graphql"
)

func GetRolesQueryField() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(roleType),
		Resolve: rolesResolve,
	}
}
