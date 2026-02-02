package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определение команд
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", 0, "ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDescription := addCmd.String("description", "", "Наименование позиции")
	addAmount := addCmd.Int("amount", 0, "Цена за позицию")

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryMonth := summaryCmd.Int("month", 0, "месяц")

	if len(os.Args) < 2 {
		fmt.Println("Неверная комманда")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		addExpense(*addDescription, *addAmount)
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		sumExpense(*summaryMonth)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		deleteExpense(*deleteId)
	case "list":
		getList()
	default:
		fmt.Println("wrong arguments")
	}
}
