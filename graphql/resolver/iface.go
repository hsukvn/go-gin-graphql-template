package resolver

import (
	"context"
	"fmt"

	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/graphql/scalar"
	"github.com/hsukvn/go-gin-graphql-template/lib/network"
)

const procNetDevPath = "/proc/net/dev"

type ifaceArgs struct {
	Name string
}

func (r *Resolver) Iface(ctx context.Context, args ifaceArgs) (*ifaceResolver, error) {
	iface, err := network.NewIface(args.Name)
	if err != nil {
		return nil, fmt.Errorf("network: Fail to new interface (%v), err: (%v)", args.Name, err)
	}

	return &ifaceResolver{
		iface: model.NewIface(iface),
	}, nil
}

func (r *Resolver) Ifaces(ctx context.Context) (*[]*ifaceResolver, error) {
	ifaces, err := network.NewIfaces()
	if err != nil {
		return nil, err
	}

	ifaceResolvers := make([]*ifaceResolver, 0)
	for i, _ := range ifaces {
		ifaceResolvers = append(ifaceResolvers, &ifaceResolver{
			iface: model.NewIface(ifaces[i]),
		})
	}

	return &ifaceResolvers, nil
}

type ifaceResolver struct {
	iface *model.Iface
}

func (r *ifaceResolver) Name() *string {
	return &r.iface.Name
}

func (r *ifaceResolver) Mac() *string {
	return &r.iface.Mac
}

func (r *ifaceResolver) Addrv4() *[]*addrResolver {
	addrs := make([]*addrResolver, len(r.iface.Addrv4))

	for i := range addrs {
		addrs[i] = &addrResolver{
			addr: r.iface.Addrv4[i],
		}
	}

	return &addrs
}

func (r *ifaceResolver) Addrv6() *[]*addrResolver {
	addrs := make([]*addrResolver, len(r.iface.Addrv6))

	for i := range addrs {
		addrs[i] = &addrResolver{
			addr: r.iface.Addrv6[i],
		}
	}

	return &addrs
}

func (r *ifaceResolver) MTU() *int32 {
	mtu := int32(r.iface.MTU)
	return &mtu
}

func (r *ifaceResolver) Rx() *scalar.Uint64 {
	rx := scalar.Uint64(r.iface.RxByte)
	return &rx
}

func (r *ifaceResolver) Tx() *scalar.Uint64 {
	tx := scalar.Uint64(r.iface.TxByte)
	return &tx
}
