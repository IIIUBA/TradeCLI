package commands

import (
	"fmt"
	"strconv"

	"github.com/piquette/finance-go/future"
	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "buy [symbol] [quantity]",
	Short: "Buy a specified quantity of a stock or futures contract",
	Long: `This command allows you to buy a specified quantity of a stock or futures contract 
by providing the symbol (e.g., CL=F) and quantity (e.g., 10). 
Example usage: trader-cli buy CL=F 20`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		quantity := args[1]

		qty, err := strconv.ParseFloat(quantity, 64)
		if err != nil || qty <= 0 {
			fmt.Println("Ошибка: количество должно быть положительным числом.")
			return
		}

		fmt.Println("Future example\n====================")
		fmt.Println()

		q, err := future.Get(symbol)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()

	},
}
