package main

import (
	"github.com/balabanovds/go-dump/cmd"
)

const file = ".LSA1_LS-update.pcap"

func main() {
	cmd.Execute()
	// if handle, err := pcap.OpenOffline(file); err != nil {
	// 	panic(err)
	// } else {
	// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	// 	for packet := range packetSource.Packets() {
	// 		ospf.HandlePacket(packet)
	// 	}
	// }

}
