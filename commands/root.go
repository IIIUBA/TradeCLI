package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(buyCmd)
}

var rootCmd = &cobra.Command{
	Use:   "TradeCLI",
	Short: "TradeCLI - это упрощенный симулятор трейдинга",
	Long: `
TradeCLI - это CLI инструмент для торговли. 
Главная цель этого инструмента - упростить торговлю фьючерсами и обучить торговлю на реальных данных.
									~by iiiuba`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
