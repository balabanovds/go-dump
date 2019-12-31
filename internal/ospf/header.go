package ospf

import (
	"encoding/binary"
	"net"
)

type ospfVersion uint16

const (
	v2 ospfVersion = 2
	v3 ospfVersion = 3
)

type ospfPacketType uint8

const (
	hello ospfPacketType = iota + 1
	dbd
	lsRequest
	lsUpdate
	lsAck
)

type OspfHeader struct {
	Version  ospfVersion
	Type     ospfPacketType
	Length   uint16
	RouterID net.IP
	AreaID   net.IP
}

func parseHeader(data []byte) OspfHeader {
	var h OspfHeader
	h.Version = ospfVersion(data[0])
	h.Type = ospfPacketType(data[1])
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.RouterID = net.IP(data[4:8])
	h.AreaID = net.IP(data[8:12])

	return h
}
