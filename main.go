package main

import (
	"FinTech/commands"
	"FinTech/data"
)

func main() {
	db := data.InitDB()
	defer db.Close()

	commands.SetDB(db)

	commands.Execute()
}
