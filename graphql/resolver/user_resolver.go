package resolver

import (
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/graphql/scalar"
)

type userResolver struct {
	user *model.User
}

func (r *userResolver) ID() *int32 {
	return &r.user.ID
}

func (r *userResolver) FirstName() *string {
	return &r.user.Firstname
}

func (r *userResolver) LastName() *string {
	return &r.user.Lastname
}

func (r *userResolver) Roles() *[]*roleResolver {
	roles := make([]*roleResolver, len(r.user.Roles))

	for i := range roles {
		roles[i] = &roleResolver{
			role: r.user.Roles[i],
		}
	}

	return &roles
}

func (r *userResolver) Deposit() *scalar.Int64 {
	d := scalar.Int64(r.user.Deposit)
	return &d
}
