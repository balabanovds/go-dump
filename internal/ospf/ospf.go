package ospf

import (
	"fmt"
	"net"
	"time"

	"github.com/balabanovds/go-dump/internal/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const ospfv2HeaderLen = 24

type packet struct {
	fullLen   int
	timestamp time.Time
	data      []byte
}

type ospfPacket struct {
	packet
	length   int
	routerID net.IP
	area     net.IP
	lsaNum   int
	content  layers.LSUpdate
}

type lsaCompatible interface {
	Dump() string
}

func handlePacket(p gopacket.Packet) error {

	// fmt.Println(p.Dump())

	var pck ospfPacket
	pck.fullLen = p.Metadata().Length
	pck.timestamp = p.Metadata().Timestamp
	pck.packet.data = p.Data()

	ipLayer := p.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 Layer detected")
	}

	ip, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return fmt.Errorf("Not an IPv4 packet TIMESTAMP: %v", pck.timestamp)
	}

	if ip.Protocol == layers.IPProtocolOSPF {
		fmt.Println("OSPF Layer detected")
		parseHeader(ip.Payload[:ospfv2HeaderLen])
	}

	ospfLayer := p.Layer(layers.LayerTypeOSPF)
	if ospfLayer != nil {
		fmt.Println("OSPF Layer detected")
		ospf, ok := ospfLayer.(*layers.OSPFv2)

		if !ok {
			return fmt.Errorf("Failed to parse OSPF at packet TIMESTAMP: %v", pck.timestamp)
		}

		pck.routerID = util.Int2Ip(ospf.RouterID)
		pck.area = util.Int2Ip(ospf.AreaID)
		pck.length = int(ospf.PacketLength)

		content := ospf.Content.(layers.LSUpdate)

		fmt.Printf("%v\n", content.NumOfLSAs)
		start := pck.fullLen - pck.length + ospfv2HeaderLen
		fmt.Printf("%v\n", p.Data()[start:])
	}

	return nil
}
