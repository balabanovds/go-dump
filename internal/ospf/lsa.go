package ospf

import (
	"encoding/binary"
	"net"
)

const lsaHeaderLength = 20

type LSType uint8

const (
	LSTypeRouter LSType = iota + 1
	LSTypeNetwork
	LSTypeSummary
	LSTypeSummaryASBR
	LSTypeASExternal
	LSTypeMulticast
	LSTypeNSSA
	LSTypeExtAttrBGP
)

type LSA struct {
	LSAge   uint16
	Options uint8
	LSType
	LinkStateID net.IP
	AdvRouter   net.IP
	SeqNumber   uint32
	Length      uint16
	Checksum    uint16
}

func parseLSA(data []byte) LSA {
	var lsa LSA
	var o offset
	lsa.LSAge = binary.BigEndian.Uint16(shiftN(data, &o, 2))
	lsa.Options = uint8(shiftOne(data, &o))
	lsa.LSType = LSType(shiftOne(data, &o))
	lsa.LinkStateID = net.IP(shiftN(data, &o, 4))
	lsa.AdvRouter = net.IP(shiftN(data, &o, 4))
	lsa.SeqNumber = binary.BigEndian.Uint32(shiftN(data, &o, 4))
	lsa.Checksum = binary.BigEndian.Uint16(shiftN(data, &o, 2))
	lsa.Length = binary.BigEndian.Uint16(shiftN(data, &o, 2))
	return lsa
}
