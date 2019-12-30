package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hostsFile string

var rootCmd = &cobra.Command{
	Use: "go-dump",
	// Short: "go-dump is a tool to parse OSPF or SNMP packets for troubleshooting",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVar(&hostsFile, "hosts", "", "Hosts file in format: IP NAME [TYPE]")
	rootCmd.MarkPersistentFlagRequired("hosts")
	rootCmd.AddCommand(ospfCmd, snmpCmd)
	ospfCmd.AddCommand(ospfFileCmd)
	snmpCmd.AddCommand(snmpFileCmd)
}
