package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveExpenses() {
	data, errMarshal := json.MarshalIndent(expenses, "", "  ")
	if errMarshal != nil {
		fmt.Println("Ошибка при преобразовании данных в JSON:", errMarshal)
		return
	}

	errMarshal = os.WriteFile(
		"expenses.json",
		data,
		0644,
	)
	if errMarshal != nil {
		fmt.Println("Ошибка при сохранении файла:", errMarshal)
	}

	fmt.Println("Данные успешно сохранены")
}

func loadExpenses() {
	fileName := "./expenses.json"

	data, errReadFile := os.ReadFile(fileName)
	if errReadFile != nil {
		if os.IsNotExist(errReadFile) {
			expenses = []Expense{}
			nextID = 1
			fmt.Println("Файл expenses.json не найден, при сохранении будет создан новый.")
			return
		} else {
			fmt.Println("Ошибка при попыткe прочесть файл: ", errReadFile)
			return
		}
	}

	errReadFile = json.Unmarshal(data, &expenses)
	if errReadFile != nil {
		fmt.Println("Ошибка при разборе JSON:", errReadFile)
		expenses = []Expense{}
		nextID = 1
		return
	}

	maxID := 0
	for _, exp := range expenses {
		if exp.ID > maxID {
			maxID = exp.ID
		}
	}

	nextID = maxID + 1

	fmt.Printf("Загружено %d расходов из файла.\n", len(expenses))
}
