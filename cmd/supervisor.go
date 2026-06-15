package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var supervisorCmd = &cobra.Command{
	Use:     "supervisor",
	Aliases: []string{"super", "su"},
	Short:   "Monitor, control and configure the Factory Assistant Supervisor",
	Long: `
The Factory Assistant Supervisor is the heart of the Factory Assistant system.
It manages your Factory Assistant Core, Operating System, and all the apps.
It even manages itself! This series of command give you control over the
Factory Assistant Supervisor.`,
	Example: `
  ha supervisor reload
  ha supervisor update
  ha supervisor logs`,
}

func init() {
	slog.Debug("Init supervisor")
	rootCmd.AddCommand(supervisorCmd)
}
