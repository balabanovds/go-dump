package packet

import (
	"fmt"
	"time"

	"github.com/balabanovds/go-dump/internal/ospf"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Packet struct {
	Length    int
	Timestamp time.Time
}

func handleIPPacket(l gopacket.Layer, p Packet) error {

	ip, ok := l.(*layers.IPv4)
	if !ok {
		return fmt.Errorf("Not an IPv4 packet TIMESTAMP: %v", p.Timestamp)
	}

	if ip.Protocol == layers.IPProtocolOSPF {
		fmt.Println("OSPF Layer detected")
		p, err := ospf.ParsePacket(ip.Payload)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", p)
		//TODO save parsed packet somewhere
	}

	return nil
}
