package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
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
	fmt.Print("Напишите ID расхода, который нужно изменить")
	scanner.Scan()
	IDinput := scanner.Text()

	ID, errID := strconv.Atoi(IDinput)
	if errID != nil {
		fmt.Println("ID должно быть числом")
		return
	}

	index := -1
	for i, exp := range expenses {
		if ID == exp.ID {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Расход с таким ID не найден")
		return
	}

	fmt.Println("Что необходимо заменить? Описание/Сумма/Дата")
	scanner.Scan()
	choice := strings.ToLower(strings.TrimSpace(scanner.Text()))

	switch choice {
	case "описание":
		fmt.Print("Введите новое описание: ")
		scanner.Scan()
		expenses[index].Description = scanner.Text()
		fmt.Println("Описание успешно обновлено!")

	case "сумма":
		fmt.Print("Введите новую сумму: ")
		scanner.Scan()
		amountInput := scanner.Text()

		newAmount, err := strconv.ParseFloat(amountInput, 64)
		if err != nil {
			fmt.Println("Ошибка: сумма должна быть числом.")
			return
		}
		expenses[index].Amount = newAmount
		fmt.Println("Сумма успешно обновлена!")

	case "дата":
		fmt.Print("Введите новую дату (в формате ГГГГ-ММ-ДД): ")
		scanner.Scan()
		dateInput := scanner.Text()

		newDate, err := time.Parse("2006-01-02", dateInput)
		if err != nil {
			fmt.Println("Ошибка: неверный формат даты. Используйте ГГГГ-ММ-ДД.")
			return
		}
		expenses[index].Date = newDate
		fmt.Println("Дата успешно обновлена!")
	default:
		fmt.Println("Неизвестный параметр. Изменения не внесены.")
	}
}

func deleteExpense(scanner *bufio.Scanner) {
	fmt.Print("Напишите ID расхода, который нужно удалить")
	scanner.Scan()
	IDinput := scanner.Text()

	ID, errID := strconv.Atoi(IDinput)
	if errID != nil {
		fmt.Println("ID должно быть числом")
		return
	}

	index := -1
	for i, exp := range expenses {
		if ID == exp.ID {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Расход с таким ID не найден")
		return
	}
	expenses = append(
		expenses[:index],
		expenses[index+1:]...,
	)

	fmt.Println("Расход успешно удалён")
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

	month, errMonth := strconv.Atoi(monthInput)
	if errMonth != nil {
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
