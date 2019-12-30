package cmd

import (
	"fmt"

	"github.com/balabanovds/go-dump/internal/ospf"

	"github.com/spf13/cobra"
)

var ospfFileCmd = &cobra.Command{
	Use:   "files",
	Short: "Parse local pcap files for OSPF packets (several packets will be merged)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ospf.RunOffline(hostsFile, args...)
	},
}

var snmpFileCmd = &cobra.Command{
	Use:   "files",
	Short: "Parse local pcap files for SNMP packets (several packets will be merged)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This option is not implemented yet")
	},
}
