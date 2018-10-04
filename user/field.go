package user

import (
	"github.com/graphql-go/graphql"
)

func UsersQueryFieldGet() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(userType),
		Resolve: usersResolveGet,
	}
}

func CreateUserMutationFieldGet() *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: createUserResolveGet,
	}
}
