package ospf

import (
	"encoding/binary"
	"fmt"
	"net"
)

const ls1LinkLength = 12

// LS1 represents OSPF Router-LSA struct
type LS1 struct {
	LSA
	LS1Flags
	NumOfLinks uint16
	Links      []LS1Link
}

// LS1Flags represent flags of Router-LSA packet
type LS1Flags struct {
	VirtualLink bool
	ASBR        bool
	ABR         bool
}

// LS1Link is one of link types
type LS1Link struct {
	LinkID     net.IP
	Data       net.IP
	Type       LS1LinkType
	MetricsNum uint8
	Metric     uint16
}

// LS1LinkType is a Router-LSA link of type:
// - PTP     where LinkID == Neighbor router ID,
// - Transit where LinkID == IP address of DR,
// - Stub    where LinkID == IP Network,
// - Virtual where LinkID == Neighbor router ID
type LS1LinkType uint16

const (
	// PTP - point to point link type
	PTP LS1LinkType = iota + 1
	// Transit link type (broadcast)
	Transit
	// Stub - connection to stub network
	Stub
	// Virtual link
	Virtual
)

func parseLS1(data []byte) LS1 {
	var l LS1
	var o offset

	f := uint8(shiftOne(data, &o))
	f8 := fmt.Sprintf("%04b", f)
	l.LS1Flags.VirtualLink = f8[1] == '1'
	l.LS1Flags.ASBR = f8[2] == '1'
	l.LS1Flags.ABR = f8[3] == '1'

	// we shift one byte coz it is '0' and used for padding?
	shiftOne(data, &o)

	l.NumOfLinks = binary.BigEndian.Uint16(shiftN(data, &o, 2))

	for i := 0; i < int(l.NumOfLinks); i++ {
		link := LS1Link{
			LinkID:     net.IP(shiftN(data, &o, 4)),
			Data:       net.IP(shiftN(data, &o, 4)),
			Type:       LS1LinkType(shiftOne(data, &o)),
			MetricsNum: uint8(shiftOne(data, &o)),
			Metric:     binary.BigEndian.Uint16(shiftN(data, &o, 2)),
		}
		l.Links = append(l.Links, link)
	}
	return l
}
