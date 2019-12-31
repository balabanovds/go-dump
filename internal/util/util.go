package util

import (
	"encoding/binary"
	"net"
)

// Offset is an util type
type Offset uint32

// ShiftOne take one byte from slice at index offset and increment offset
func ShiftOne(data []byte, off *Offset) (b byte) {
	b = data[*off]
	*off++
	return
}

// ShiftN take N bytes from slice starting at index offset and increment offset by N
func ShiftN(data []byte, off *Offset, n int) (b []byte) {
	b = data[*off : *off+Offset(n)]
	*off += Offset(n)
	return
}

// Int2Ip converts uint32 pcap representation to IP
func Int2Ip(i uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, i)
	return ip
}
