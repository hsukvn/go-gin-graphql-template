package user

import (
	"github.com/graphql-go/graphql"
)

func usersResolveGet(p graphql.ResolveParams) (interface{}, error) {
	return users, nil
}

func createUserResolveGet(p graphql.ResolveParams) (interface{}, error) {
	u := user{
		ID:        len(users)+1,
		Firstname: p.Args["firstname"].(string),
		Lastname:  p.Args["lastname"].(string),
	}

	// FIXME: Mockup
	users = append(users, u)

	return u, nil
}
