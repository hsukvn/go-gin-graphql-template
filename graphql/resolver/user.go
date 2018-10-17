package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/lib/user"
)

type userArgs struct {
	UID string
}

func (r *Resolver) User(ctx context.Context, args userArgs) (*userResolver, error) {
	user, err := model.NewUserByUID(args.UID)
	if err != nil {
		return nil, fmt.Errorf("user: Fail to new user (%v), err: (%v)", args.UID, err)
	}

	return &userResolver{
		user: user,
	}, nil
}

type userByNameArgs struct {
	Name string
}

func (r *Resolver) UserByName(ctx context.Context, args userByNameArgs) (*userResolver, error) {
	user, err := model.NewUserByName(args.Name)
	if err != nil {
		return nil, fmt.Errorf("user: Fail to new user (%v), err: (%v)", args.Name, err)
	}

	return &userResolver{
		user: user,
	}, nil
}

func (r *Resolver) Users(ctx context.Context) (*[]*userResolver, error) {
	users, err := user.ListUsers()
	if err != nil {
		return nil, fmt.Errorf("user: Fail to list users, err: (%v)", err)
	}

	userResolvers := make([]*userResolver, 0)
	for _, user := range users {
		u, err := model.NewUserByName(user)
		if err != nil {
			continue
		}

		userResolvers = append(userResolvers, &userResolver{
			user: u,
		})
	}

	return &userResolvers, nil
}

type userResolver struct {
	user *model.User
}

func (r *userResolver) UID() *graphql.ID {
	uid := graphql.ID(r.user.UID)
	return &uid
}

func (r *userResolver) GID() *graphql.ID {
	gid := graphql.ID(r.user.GID)
	return &gid
}

func (r *userResolver) Name() *string {
	return &r.user.Name
}

func (r *userResolver) Home() *string {
	return &r.user.Home
}

func (r *userResolver) Groups() *[]*groupResolver {
	groups := make([]*groupResolver, len(r.user.Groups))

	for i := range groups {
		groups[i] = &groupResolver{
			group: r.user.Groups[i],
		}
	}

	return &groups
}
