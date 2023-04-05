package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xomander/mandery/constant"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mandery",
	Run:   printVersion,
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion(cmd *cobra.Command, args []string) {
	os.Stdout.WriteString("mandery version " + constant.Version + "\n")
}
