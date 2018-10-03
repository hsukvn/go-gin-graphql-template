package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/hsukvn/go-graphql-template/types"
)

func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := types.User{
				ID: len(types.Users)+1,
				Firstname: p.Args["firstname"].(string),
				Lastname:  p.Args["lastname"].(string),
			}

			// FIXME: Mockup
			types.Users = append(types.Users, user)

			return user, nil
		},
	}
}
