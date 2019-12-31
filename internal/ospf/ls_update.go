package ospf

import (
	"encoding/binary"

	u "github.com/balabanovds/go-dump/internal/util"
)

// LSUpdate packet struct
type LSUpdate struct {
	NumOfLSAa  uint32
	LS1Packets []LS1
	LS2Packets []LS2
}

type offset uint32

func parseLSUpdatePacket(data []byte) LSUpdate {
	var ls LSUpdate

	var off u.Offset

	ls.NumOfLSAa = binary.BigEndian.Uint32(u.ShiftN(data, &off, 4))

	for i := 0; i < int(ls.NumOfLSAa); i++ {
		length := binary.BigEndian.Uint16(data[off+18 : off+20])
		end := off + u.Offset(length)
		lsa := parseLSA(data[off:end])

		switch lsa.LSType {
		case LSTypeRouter:
			ls1 := parseLS1(data[off+lsaHeaderLength : end])
			ls1.LSA = lsa
			ls.LS1Packets = append(ls.LS1Packets, ls1)
		case LSTypeNetwork:
			ls2 := parseLS2(data[off+lsaHeaderLength : end])
			ls2.LSA = lsa
			ls.LS2Packets = append(ls.LS2Packets, ls2)
		}

		off = end
	}

	return ls
}
