package ospf

// Packet - full OSPF packet
type Packet struct {
	Header
	LSUpdate
	// TODO: add other types here, now actual only LS-Update
}

// ParsePacket parsing OSPF data payload
func ParsePacket(data []byte) (Packet, error) {
	var p Packet

	p.Header = parseHeader(data[:ospfv2HeaderLen])

	//TODO add parsing others types
	switch p.Header.Type {
	case lsUpdate:
		p.LSUpdate = parseLSUpdatePacket(data[ospfv2HeaderLen:])
	}

	return p, nil
}
