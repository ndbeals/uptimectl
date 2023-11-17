package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ndbeals/uptimectl/cmd/auth"
	"github.com/ndbeals/uptimectl/cmd/config"
	"github.com/ndbeals/uptimectl/cmd/incidents"
	"github.com/ndbeals/uptimectl/cmd/monitorgroups"
	"github.com/ndbeals/uptimectl/cmd/monitors"
	"github.com/ndbeals/uptimectl/cmd/oncall"
	"github.com/ndbeals/uptimectl/cmd/statuspages"
	"github.com/ndbeals/uptimectl/cmd/version"
	"github.com/ndbeals/uptimectl/pkg/authmanager"
	"github.com/ndbeals/uptimectl/pkg/contextmanager"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "uptimectl",
	Short: "A command-line interface for working with Better Uptime",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		switch err {
		case authmanager.ErrNoLogin:
			fallthrough
		case authmanager.ErrSessionNotActive:
			_, _ = fmt.Fprintf(os.Stderr, "%v, please use: uptimectl auth login\n", err)
		default:
			_, _ = fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(auth.AuthCmd)
	RootCmd.AddCommand(incidents.IncidentsCmd)
	RootCmd.AddCommand(monitorgroups.MonitorGroupsCmd)
	RootCmd.AddCommand(statuspages.StatusPagesCmd)
	RootCmd.AddCommand(config.ConfigCmd)
	RootCmd.AddCommand(version.VersionCmd)
	RootCmd.AddCommand(oncall.OncallCmd)
	RootCmd.AddCommand(monitors.MonitorsCmd)
	cobra.OnInitialize(contextmanager.Init)
}
