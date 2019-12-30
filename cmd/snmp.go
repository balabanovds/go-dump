package cmd

import (
	"github.com/spf13/cobra"
)

var snmpCmd = &cobra.Command{
	Use:   "snmp",
	Short: "parse network interface or pcap file for SNMP packets",
	Args:  cobra.MinimumNArgs(1),
}
