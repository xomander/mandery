package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xomander/mandery/constant"
	"os"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run mandery",
	Run:   run,
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	os.Stdout.WriteString("mandery version " + constant.Version + "\n")
}
