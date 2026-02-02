package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id         int       `json:"id"`
	Dscription string    `json:"description"`
	Date       time.Time `json:"createdAt"`
	Amount     int       `json:"updatedAt"`
}

func createNewExpense(s string, amount int) Expense {
	expense := Expense{
		Id:         0,
		Dscription: s,
		Date:       time.Now().UTC(),
		Amount:     amount,
	}
	return expense
}

func addExpense(s string, amount int) {
	expense := createNewExpense(s, amount)

	_, expenses := readFile()

	expense.Id = getNextId(&expenses)

	expenses = append(expenses, expense)

	writeFile(&expenses)

	fmt.Printf("Expense added successfully (ID: %d)", expense.Id)
}

func getNextId(expenses *[]Expense) int {
	maxId := 0
	for _, expense := range *expenses {
		if expense.Id > maxId {
			maxId = expense.Id
		}
	}
	return maxId + 1
}

func findExpenseIdxById(id int, expenses *[]Expense) (bool, int) {
	for idx, expense := range *expenses {
		if expense.Id == id {
			return true, idx
		}
	}

	return false, 0
}

func sumExpense(month int) {
	_, expenses := readFile()

	summuryExpenses := 0

	for _, exexpense := range expenses {
		if month == 0 {
			summuryExpenses += exexpense.Amount
		} else if int(exexpense.Date.Month()) == month {
			summuryExpenses += exexpense.Amount
		}
	}

	fmt.Printf("Total expenses: $%d\n", summuryExpenses)
}

func deleteExpense(id int) {
	_, expenses := readFile()

	ok, idx := findExpenseIdxById(id, &expenses)

	if ok {
		expenses = append(expenses[:idx], expenses[idx+1:]...)
		writeFile(&expenses)
		return
	}

	fmt.Println("Нет элемента с Id", id)
}

func getList() {
	_, expenses := readFile()
	fmt.Println("ID  Date       Description  Amount")

	for _, exexpense := range expenses {
		fmt.Printf("%d %s %s $%d\n", exexpense.Id, exexpense.Date.Format("2006-01-02"), exexpense.Dscription, exexpense.Amount)
	}
}
