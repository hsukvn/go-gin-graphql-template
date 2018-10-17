package resolver

import (
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
)

type addrResolver struct {
	addr *model.Address
}

func (r *addrResolver) IP() *string {
	return &r.addr.IP
}

func (r *addrResolver) Mask() *string {
	return &r.addr.Mask
}
