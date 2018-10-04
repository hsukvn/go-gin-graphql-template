package role

import (
	"github.com/graphql-go/graphql"
)

func rolesResolveGet(p graphql.ResolveParams) (interface{}, error) {
	return roles, nil
}
