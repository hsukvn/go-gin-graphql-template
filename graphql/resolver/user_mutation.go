package resolver

import (
	"context"

	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
)

type addUserArgs struct {
	Firstname string
	Lastname  string
	Roles     []*string
}

func (r *Resolver) AddUser(ctx context.Context, args addUserArgs) (*userResolver, error) {
	roles := make([]*model.Role, 0)

	for i, r := range args.Roles {
		roles = append(roles, &model.Role{
			ID:   int32(len(model.UserData)) + int32(i) + 1,
			Name: *r,
		})
	}

	u := userResolver{
		user: &model.User{
			ID:        int32(len(model.UserData)) + 1,
			Firstname: args.Firstname,
			Lastname:  args.Lastname,
			Roles:     roles,
			Deposit:   0,
		},
	}

	return &u, nil
}
