package main

import (
	"bufio"
	"fmt"
	"strconv"
	"time"
)

func addExpense(scanner *bufio.Scanner) {
	fmt.Print("Добавьте описание: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Println("Введите сумму: ")

	scanner.Scan()
	amountInput := scanner.Text()
	amount, errFloat := strconv.ParseFloat(amountInput, 64)
	if errFloat != nil {
		fmt.Println("Ошибка! Сумма должна быть числом")
		return
	}

	expense := Expense{
		ID:          nextID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	nextID++

	expenses = append(expenses, expense)

	fmt.Println("Вы успешно добавили расход")
}

func updateExpense(scanner *bufio.Scanner) {
	// TODO:
	// 1. Запросить ID
	// 2. Найти расход
	// 3. Изменить данные
}

func deleteEpxense(scanner *bufio.Scanner) {
	// TODO:
	// 1. Запросить ID
	// 2. Найти запись
	// 3. Удалить из slice
}

func listExpenses() {

	if len(expenses) == 0 {
		fmt.Println("Расходов не обнаружено")
		return
	}

	fmt.Printf("%-5s | %-20s | %-10s | %-10s\n",
		"ID",
		"DESCRIPTION",
		"AMOUNT",
		"DATE",
	)

	fmt.Println("-------------------------------------------------------------")

	for _, exp := range expenses {
		fmt.Printf("%-5d | %-20s | %-10.2f | %-12s\n",
			exp.ID,
			exp.Description,
			exp.Amount,
			exp.Date.Format("2006-01-02"),
		)
	}
}

func showMonthSummary(scanner *bufio.Scanner) {
	total := 0.0

	fmt.Print("Напишите номер месяца (1-12): ")

	scanner.Scan()
	monthInput := scanner.Text()

	month, err := strconv.Atoi(monthInput)
	if err != nil {
		fmt.Println("Ошибка: введите число от 1 до 12")
		return
	}

	if month < 1 || month > 12 {
		fmt.Println("Ошибка: месяц должен быть от 1 до 12! ")
		return
	}

	currentYear := time.Now().Year()

	for _, exp := range expenses {
		if exp.Date.Month() == time.Month(month) &&
			exp.Date.Year() == currentYear {

			total += exp.Amount
		}
	}

	fmt.Printf(
		"Общая сумма трат за %s %d года составила: %.2f\n",
		time.Month(month),
		currentYear,
		total,
	)
}

func exportCSV() {
	// TODO:
	// 1. Создать CSV файл
	// 2. Записать данные
}
