package network

import (
	"fmt"
	"net"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Address struct {
	IP   string
	Mask string
}

type Iface struct {
	Name   string
	Mac    string
	Addrv4 []*Address
	Addrv6 []*Address
	MTU    int
	Stat   *linuxproc.NetworkStat
}

const procNetDevPath = "/proc/net/dev"

func NewIface(name string) (*Iface, error) {
	stats, err := getNetworkStats()
	if err != nil {
		return nil, err
	}

	if _, ok := stats[name]; !ok {
		return nil, fmt.Errorf("interface %v does not exist", name)
	}

	iface, err := newIfaceByStat(stats[name])
	if err != nil {
		return nil, err
	}

	return iface, nil
}

func NewIfaces() ([]*Iface, error) {
	stats, err := getNetworkStats()
	if err != nil {
		return nil, err
	}

	ifaces := make([]*Iface, 0)

	for name, _ := range stats {
		iface, err := newIfaceByStat(stats[name])
		if err != nil {
			continue
		}
		ifaces = append(ifaces, iface)
	}

	return ifaces, nil
}

func getNetworkStats() (map[string]*linuxproc.NetworkStat, error) {
	stats, err := linuxproc.ReadNetworkStat(procNetDevPath)
	if err != nil {
		return nil, err
	}

	networkStats := make(map[string]*linuxproc.NetworkStat)

	for i, s := range stats {
		networkStats[s.Iface] = &stats[i]
	}

	return networkStats, nil
}

func newIfaceByStat(s *linuxproc.NetworkStat) (*Iface, error) {
	iface, err := net.InterfaceByName(s.Iface)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	addrv4 := make([]*Address, 0)
	addrv6 := make([]*Address, 0)

	for _, addr := range addrs {
		var ip net.IP
		var mask net.IPMask

		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
			mask = v.Mask
		case *net.IPAddr:
			ip = v.IP
			mask = ip.DefaultMask()
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		if ip.To4() != nil {
			addrv4 = append(addrv4, &Address{
				IP:   ip.String(),
				Mask: ipv4MaskString(mask),
			})
		} else {
			addrv6 = append(addrv6, &Address{
				IP:   ip.String(),
				Mask: ipv6MaskString(mask),
			})
		}
	}

	return &Iface{
		Name:   s.Iface,
		Mac:    iface.HardwareAddr.String(),
		Addrv4: addrv4,
		Addrv6: addrv6,
		MTU:    iface.MTU,
		Stat:   s,
	}, nil
}

func ipv4MaskString(mask net.IPMask) string {
	const maxIPv4StringLen = len("255.255.255.255")
	b := make([]byte, maxIPv4StringLen)

	n := ubtoa(b, 0, mask[0])
	b[n] = '.'
	n++
	n += ubtoa(b, n, mask[1])
	b[n] = '.'
	n++
	n += ubtoa(b, n, mask[2])
	b[n] = '.'
	n++
	n += ubtoa(b, n, mask[3])

	return string(b[:n])
}

func ipv6MaskString(mask net.IPMask) string {
	const (
		IPv6len = 16
		maxLen  = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	)

	p := mask

	// Find longest run of zeros.
	e0 := -1
	e1 := -1
	for i := 0; i < IPv6len; i += 2 {
		j := i
		for j < IPv6len && p[j] == 0 && p[j+1] == 0 {
			j += 2
		}
		if j > i && j-i > e1-e0 {
			e0 = i
			e1 = j
			i = j
		}
	}
	// The symbol "::" MUST NOT be used to shorten just one 16 bit 0 field.
	if e1-e0 <= 2 {
		e0 = -1
		e1 = -1
	}

	b := make([]byte, 0, maxLen)

	// Print with possible :: in place of run of zeros
	for i := 0; i < IPv6len; i += 2 {
		if i == e0 {
			b = append(b, ':', ':')
			i = e1
			if i >= IPv6len {
				break
			}
		} else if i > 0 {
			b = append(b, ':')
		}
		b = appendHex(b, (uint32(p[i])<<8)|uint32(p[i+1]))
	}
	return string(b)
}

func ubtoa(dst []byte, start int, v byte) int {
	if v < 10 {
		dst[start] = v + '0'
		return 1
	} else if v < 100 {
		dst[start+1] = v%10 + '0'
		dst[start] = v/10 + '0'
		return 2
	}

	dst[start+2] = v%10 + '0'
	dst[start+1] = (v/10)%10 + '0'
	dst[start] = v/100 + '0'
	return 3
}

func appendHex(dst []byte, i uint32) []byte {
	const hexDigit = "0123456789abcdef"

	if i == 0 {
		return append(dst, '0')
	}
	for j := 7; j >= 0; j-- {
		v := i >> uint(j*4)
		if v > 0 {
			dst = append(dst, hexDigit[v&0xf])
		}
	}
	return dst
}
