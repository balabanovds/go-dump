package ospf

import (
	"encoding/binary"
	"net"
)

// Version - OSPF version
type Version uint16

const (
	v2 Version = iota + 2
	v3
)

const ospfv2HeaderLen = 24

// PacketType - OSPF packet type
type PacketType uint8

const (
	hello PacketType = iota + 1
	dbd
	lsRequest
	lsUpdate
	lsAck
)

// Header - OSPF header
type Header struct {
	Version  Version
	Type     PacketType
	Length   uint16
	RouterID net.IP
	AreaID   net.IP
}

func parseHeader(data []byte) Header {
	var h Header
	h.Version = Version(data[0])
	h.Type = PacketType(data[1])
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.RouterID = net.IP(data[4:8])
	h.AreaID = net.IP(data[8:12])

	return h
}
