package ospf

import (
	"net"
)

// LS2 represents OSPF Network-LSA struct
type LS2 struct {
	LSA
	Netmask         net.IP
	AttachedRouters []net.IP
}

func parseLS2(data []byte) LS2 {
	var l LS2
	var o offset

	l.Netmask = net.IP(shiftN(data, &o, 4))

	for {
		if len(data) == int(o) {
			break
		}
		l.AttachedRouters = append(l.AttachedRouters, net.IP(shiftN(data, &o, 4)))
	}

	return l
}
