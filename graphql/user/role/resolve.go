package role

import (
	"github.com/graphql-go/graphql"
)

func rolesResolve(p graphql.ResolveParams) (interface{}, error) {
	return roles, nil
}
