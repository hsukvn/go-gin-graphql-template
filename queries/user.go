package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/hsukvn/go-graphql-template/types"
)

func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			users := types.Users
			return users, nil
		},
	}
}
