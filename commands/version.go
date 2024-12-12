package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TradeCLI",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TradeCLI trader terminal v0.1.1 -- Shiba")
	},
}
