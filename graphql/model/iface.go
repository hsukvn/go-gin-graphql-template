package model

import (
	"github.com/hsukvn/go-gin-graphql-template/lib/network"
)

type Iface struct {
	Name   string     `json:"name"`
	Mac    string     `json:"mac"`
	Addrv4 []*Address `json:"addrv4"`
	Addrv6 []*Address `json:"addrv6"`
	MTU    int        `json:"mtu"`
	RxByte uint64     `json:"rxbyte"`
	TxByte uint64     `json:"txbyte"`
}

func NewIface(iface *network.Iface) *Iface {
	return &Iface{
		Name:   iface.Name,
		Mac:    iface.Mac,
		Addrv4: getAddrs(iface.Addrv4),
		Addrv6: getAddrs(iface.Addrv6),
		MTU:    iface.MTU,
		RxByte: iface.Stat.RxBytes,
		TxByte: iface.Stat.TxBytes,
	}
}

func getAddrs(addrs []*network.Address) []*Address {
	modelAddrs := make([]*Address, len(addrs))

	for i, _ := range modelAddrs {
		modelAddrs[i] = &Address{
			IP:   addrs[i].IP,
			Mask: addrs[i].Mask,
		}
	}

	return modelAddrs
}
