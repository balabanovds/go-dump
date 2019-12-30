package util

import (
	"encoding/binary"
	"net"
)

// Int2Ip converts uint32 pcap representation to IP
func Int2Ip(i uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, i)
	return ip
}
