package cmd

import (
	"github.com/spf13/cobra"
)

var multicastCmd = &cobra.Command{
	Use:     "multicast",
	Aliases: []string{"mcast", "mc"},
	Short:   "Get information, update or configure the Factory Assistant Multicast",
	Long: `
The multicast command allows you to manage the internal Factory Assistant Multicast
backend by exposing commands to view, monitor, configure and control it.`,
	Example: `
  ha multicast info
  ha multicast update`,
}

func init() {
	rootCmd.AddCommand(multicastCmd)
}
