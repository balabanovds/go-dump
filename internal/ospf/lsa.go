package ospf

import (
	"encoding/binary"
	"net"

	u "github.com/balabanovds/go-dump/internal/util"
)

const lsaHeaderLength = 20

// LSType ..
type LSType uint8

const (
	// LSTypeRouter LSA1
	LSTypeRouter LSType = iota + 1
	// LSTypeNetwork LSA2
	LSTypeNetwork
	// LSTypeSummary LSA3
	LSTypeSummary
	// LSTypeSummaryASBR LSA4
	LSTypeSummaryASBR
	// LSTypeASExternal LSA5
	LSTypeASExternal
	// LSTypeMulticast LSA6
	LSTypeMulticast
	// LSTypeNSSA LSA7
	LSTypeNSSA
	// LSTypeExtAttrBGP LSA8
	LSTypeExtAttrBGP
)

// LSA is a base LSA struct
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
	var o u.Offset
	lsa.LSAge = binary.BigEndian.Uint16(u.ShiftN(data, &o, 2))
	lsa.Options = uint8(u.ShiftOne(data, &o))
	lsa.LSType = LSType(u.ShiftOne(data, &o))
	lsa.LinkStateID = net.IP(u.ShiftN(data, &o, 4))
	lsa.AdvRouter = net.IP(u.ShiftN(data, &o, 4))
	lsa.SeqNumber = binary.BigEndian.Uint32(u.ShiftN(data, &o, 4))
	lsa.Checksum = binary.BigEndian.Uint16(u.ShiftN(data, &o, 2))
	lsa.Length = binary.BigEndian.Uint16(u.ShiftN(data, &o, 2))
	return lsa
}
