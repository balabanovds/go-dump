package cmd

import (
	"github.com/spf13/cobra"
)

var ospfCmd = &cobra.Command{
	Use:   "ospf",
	Short: "parse network interace or pcap file for OSPF packets",
	Args:  cobra.MinimumNArgs(1),
}
