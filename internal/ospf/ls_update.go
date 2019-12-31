package ospf

import (
	"encoding/binary"
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

	var off offset

	ls.NumOfLSAa = binary.BigEndian.Uint32(shiftN(data, &off, 4))

	for i := 0; i < int(ls.NumOfLSAa); i++ {
		length := binary.BigEndian.Uint16(data[off+18 : off+20])
		end := off + offset(length)
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

func shiftOne(data []byte, off *offset) (b byte) {
	b = data[*off]
	*off++
	return
}

func shiftN(data []byte, off *offset, n int) (b []byte) {
	b = data[*off : *off+offset(n)]
	*off += offset(n)
	return
}
