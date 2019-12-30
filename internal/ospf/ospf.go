package ospf

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/balabanovds/go-dump/internal/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type packet struct {
	length    int
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

func handlePacket(p gopacket.Packet) {

	fmt.Println(p.Dump())

	var pck ospfPacket
	pck.packet.length = p.Metadata().Length
	pck.timestamp = p.Metadata().Timestamp
	pck.packet.data = p.Data()

	ospfLayer := p.Layer(layers.LayerTypeOSPF)
	if ospfLayer != nil {
		fmt.Println("OSPF Layer detected")
		ospf, ok := ospfLayer.(*layers.OSPFv2)

		if !ok {
			log.Printf("Failed to parse OSPFv2 packet")
			return
		}

		pck.routerID = util.Int2Ip(ospf.RouterID)
		pck.area = util.Int2Ip(ospf.AreaID)
		pck.length = int(ospf.PacketLength)

		fmt.Printf("%v\n", util.Int2Ip(ospf.RouterID))
	}
}
