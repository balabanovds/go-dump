package ospf



// Packet - full OSPF packet
type Packet struct {
	Header
	LsaNum int
	LSUpdate
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

	// ospfLayer := p.Layer(layers.LayerTypeOSPF)
	// if ospfLayer != nil {
	// 	fmt.Println("OSPF Layer detected")
	// 	ospf, ok := ospfLayer.(*layers.OSPFv2)

	// 	if !ok {
	// 		return fmt.Errorf("Failed to parse OSPF at packet TIMESTAMP: %v", pck.timestamp)
	// 	}

	// 	pck.routerID = util.Int2Ip(ospf.RouterID)
	// 	pck.area = util.Int2Ip(ospf.AreaID)
	// 	pck.length = int(ospf.PacketLength)

	// 	content := ospf.Content.(layers.LSUpdate)

	// 	fmt.Printf("%v\n", content.NumOfLSAs)
	// 	start := pck.fullLen - pck.length + ospfv2HeaderLen
	// 	fmt.Printf("%v\n", p.Data()[start:])
	// }

	return p, nil
}
