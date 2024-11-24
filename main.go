package main

import (
	"FinTech/commands"
)

func main() {

	// db, err := sql.Open("sqlite", "database.db")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	commands.Execute()

}
