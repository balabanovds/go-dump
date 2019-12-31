package ospf

import (
	"encoding/binary"
	"fmt"
	"net"
)

const ls1LinkLength = 12

type LS1 struct {
	LSA
	LS1Flags
	NumOfLinks uint16
	Links      []LS1Link
}

type LS1Flags struct {
	VirtualLink bool
	ASBR        bool
	ABR         bool
}

type LS1Link struct {
	LinkID     net.IP
	Data       net.IP
	Type       LS1LinkType
	MetricsNum uint8
	Metric     uint16
}

type LS1LinkType uint16

const (
	PTP     LS1LinkType = iota + 1 // LinkID == Neighbor router ID
	Transit                        // LinkID == IP address of DR
	Stub                           // LinkID == IP Network
	Virtual                        // LinkID == Neighbor router ID
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
