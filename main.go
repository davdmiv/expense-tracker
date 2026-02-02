package main

import (
	"fmt"
	"os"

	// "strconv"
	"flag"
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
		fmt.Printf("Получение addDescription: %s addAmount: %d ", *addDescription, *addAmount)
		addExpense(*addDescription, *addAmount)
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		fmt.Printf("Получение summaryMonth: %d ", *summaryMonth)

		sumExpense(*summaryMonth)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		fmt.Printf("Получение deleteId: %d ", *deleteId)

		deleteExpense(*deleteId)
	case "list":
		getList()
	default:
		fmt.Println("wrong arguments")
	}
}
