package user

import (
	"github.com/graphql-go/graphql"
)

func usersResolve(p graphql.ResolveParams) (interface{}, error) {
	return users, nil
}

func createUserResolve(p graphql.ResolveParams) (interface{}, error) {
	u := user{
		ID:        len(users)+1,
		Firstname: p.Args["firstname"].(string),
		Lastname:  p.Args["lastname"].(string),
	}

	// FIXME: Mockup
	users = append(users, u)

	return u, nil
}
