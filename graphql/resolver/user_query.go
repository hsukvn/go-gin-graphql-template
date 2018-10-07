package resolver

import (
	"context"
	"fmt"

	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
)

type userArgs struct {
	ID int32
}

func (r *Resolver) User(ctx context.Context, args userArgs) (*userResolver, error) {
	var user userResolver

	for _, u := range model.UserData {
		if args.ID == u.ID {
			user.user = &u
			break
		}
	}

	if user.user == nil {
		return nil, fmt.Errorf("user: (ID:%v) is not exist", args.ID)
	}

	return &user, nil
}

func (r *Resolver) Users(ctx context.Context) (*[]*userResolver, error) {
	users := make([]*userResolver, 0)

	for _, u := range model.UserData {
		users = append(users, &userResolver{
			user: &u,
		})
	}

	return &users, nil
}
