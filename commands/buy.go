package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "купить [символ] [кол-во]",
	Short: "Купить фьючерс",
	Long: `Эта команда позволяет вам покупать указанное количество акций или контрактов фьючерса.
Пример использования: trader-cli buy SBERF 20`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		quantity := args[1]

		qty, err := strconv.ParseFloat(quantity, 64)
		if err != nil || qty <= 0 {
			fmt.Println("Ошибка: количество должно быть положительным числом.")
			return
		}

		fmt.Println("Вы купили " + quantity + " в количестве " + symbol)

	},
}
