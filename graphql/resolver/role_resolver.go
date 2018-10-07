package resolver

import (
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() *int32 {
	return &r.role.ID
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
