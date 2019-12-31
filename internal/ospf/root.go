package ospf

import (
	"log"

	"github.com/google/gopacket"
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
