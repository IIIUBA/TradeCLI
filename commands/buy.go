package commands

import (
	"FinTech/background"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var balance float64
var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Купить фьючерс",
	Long: `Эта команда позволяет вам покупать указанное количество акций или контрактов фьючерса.
Пример использования: trader-cli buy SBERF 2`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]
		quantity := args[1]

		qty, err := strconv.ParseFloat(quantity, 64)
		if err != nil || qty <= 0 {
			fmt.Println("Ошибка: количество должно быть положительным числом.")
			return
		}

		price := background.GetStockPrice(symbol)

		totalCost := price * qty

		err = db.QueryRow("SELECT balance FROM users WHERE id = ?", 1).Scan(&balance)
		if err != nil {
			fmt.Println("Ошибка: не удалось получить баланс пользователя.")
			return
		}

		if balance < totalCost {
			fmt.Printf("Недостаточно средств. Ваш баланс: %.2f, нужно: %.2f\n", balance, totalCost)
			return
		}

		newBalance := balance - totalCost
		_, err = db.Exec("UPDATE users SET balance = ? WHERE id = ?", newBalance, 1)
		if err != nil {
			fmt.Println("Ошибка: не удалось обновить баланс пользователя.")
		}

		_, err = db.Exec("INSERT INTO transactions (user_id, symbol, quanity, price, type) VALUES (?, ?, ?, ?, ?)", 1, symbol, qty, price, "buy")
		if err != nil {
			fmt.Println("Ошибка: не удалось добавить транзакцию.")
		}

		fmt.Println("Вы купили " + quantity + " в количестве " + symbol)

	},
}
