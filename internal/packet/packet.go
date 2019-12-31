package packet

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	pcap "github.com/google/gopacket/pcap"
)

// RunOffline fires up OSPF parsing
func RunOffline(hostFile string, files ...string) {
	for _, file := range files {
		runOne(file)
	}
}

func runOne(file string) error {
	handle, err := pcap.OpenOffline(file)
	if err != nil {
		return err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		err := handlePacket(packet)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
func handlePacket(p gopacket.Packet) error {

	var pck Packet

	pck.Length = p.Metadata().Length
	pck.Timestamp = p.Metadata().Timestamp

	ipLayer := p.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 Layer detected")
		return handleIPPacket(ipLayer, pck)
	}

	return fmt.Errorf("Not known packet type. %v", p.Dump())
}
