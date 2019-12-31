package ospf

import (
	"net"
)

type LS2 struct {
	LSA
	Netmask         net.IP
	AttachedRouters []net.IP
}
