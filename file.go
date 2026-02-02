package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	EXPENSES_JSON_FILE_NAME = "expenses.json"
)

func readFile() (bool, []Expense) {
	// 1. Читаем файл
	file, err := os.ReadFile(EXPENSES_JSON_FILE_NAME)
	if err != nil {
		fmt.Println("Не удалось прочиать файл с задачами", err)
		return false, []Expense{}
	}

	// 2. Создаем переменную-срез для хранения данных
	var expenses []Expense

	// 3. Десериализуем JSON в массив структур
	err = json.Unmarshal(file, &expenses)
	if err != nil {
		fmt.Println("Не удалось десерилизовать задачи из файла", err)
		return false, []Expense{}
	}

	return true, expenses
}

func writeFile(expenses *[]Expense) {
	// 1. Создание/перезапись файла (os.O_TRUNC очищает файл перед записью)
	file, err := os.OpenFile(EXPENSES_JSON_FILE_NAME, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// 2. Сериализация в JSON с отступами (MarshalIndent)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(expenses); err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println("Данные успешно перезаписаны в " + EXPENSES_JSON_FILE_NAME)
}
