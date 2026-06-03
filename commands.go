package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func findExpenseIndexByID(id int) int {
	for i, exp := range expenses {
		if exp.ID == id {
			return i
		}
	}

	return -1
}

func addExpense(scanner *bufio.Scanner) {
	fmt.Print("Добавьте описание: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Print("Введите сумму: ")

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

	saveExpenses()

	fmt.Println("Вы успешно добавили расход")
}

func updateExpense(scanner *bufio.Scanner) {
	fmt.Print("Напишите ID расхода, который нужно изменить: ")

	scanner.Scan()
	IDinput := scanner.Text()

	ID, errID := strconv.Atoi(IDinput)
	if errID != nil {
		fmt.Println("ID должно быть числом")
		return
	}

	index := findExpenseIndexByID(ID)

	if index == -1 {
		fmt.Println("Расход с таким ID не найден")
		return
	}

	fmt.Print("Что необходимо заменить? Описание/Сумма/Дата: ")

	scanner.Scan()
	choice := strings.ToLower(strings.TrimSpace(scanner.Text()))

	switch choice {

	case "описание":
		fmt.Print("Введите новое описание: ")

		scanner.Scan()
		expenses[index].Description = scanner.Text()

		saveExpenses()

		fmt.Println("Описание успешно обновлено!")

	case "сумма":
		fmt.Print("Введите новую сумму: ")

		scanner.Scan()
		amountInput := scanner.Text()

		newAmount, err := strconv.ParseFloat(amountInput, 64)
		if err != nil {
			fmt.Println("Ошибка: сумма должна быть числом")
			return
		}

		expenses[index].Amount = newAmount

		saveExpenses()

		fmt.Println("Сумма успешно обновлена!")

	case "дата":
		fmt.Print("Введите новую дату (в формате ГГГГ-ММ-ДД): ")

		scanner.Scan()
		dateInput := scanner.Text()

		newDate, err := time.Parse("2006-01-02", dateInput)
		if err != nil {
			fmt.Println("Ошибка: неверный формат даты. Используйте ГГГГ-ММ-ДД")
			return
		}

		expenses[index].Date = newDate

		saveExpenses()

		fmt.Println("Дата успешно обновлена!")

	default:
		fmt.Println("Неизвестный параметр. Изменения не внесены.")
	}
}

func deleteExpense(scanner *bufio.Scanner) {
	fmt.Print("Напишите ID расхода, который нужно удалить: ")

	scanner.Scan()
	IDinput := scanner.Text()

	ID, errID := strconv.Atoi(IDinput)
	if errID != nil {
		fmt.Println("ID должно быть числом")
		return
	}

	index := findExpenseIndexByID(ID)

	if index == -1 {
		fmt.Println("Расход с таким ID не найден")
		return
	}

	expenses = append(
		expenses[:index],
		expenses[index+1:]...,
	)

	saveExpenses()

	fmt.Println("Расход успешно удалён")
}

func listExpenses() {
	if len(expenses) == 0 {
		fmt.Println("Расходов не обнаружено")
		return
	}

	fmt.Printf(
		"%-5s | %-20s | %-10s | %-12s\n",
		"ID",
		"DESCRIPTION",
		"AMOUNT",
		"DATE",
	)

	fmt.Println("----------------------------------------------------------------")

	for _, exp := range expenses {
		fmt.Printf(
			"%-5d | %-20s | %-10.2f | %-12s\n",
			exp.ID,
			exp.Description,
			exp.Amount,
			exp.Date.Format("2006-01-02"),
		)
	}
}

func showSummary() {
	if len(expenses) == 0 {
		fmt.Println("Расходов пока нет")
		return
	}

	total := 0.0

	for _, exp := range expenses {
		total += exp.Amount
	}

	fmt.Printf(
		"Общая сумма расходов за всё время составила: %.2f\n",
		total,
	)
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
		fmt.Println("Ошибка: месяц должен быть от 1 до 12")
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

func exportCSV(scanner *bufio.Scanner) {
	if len(expenses) == 0 {
		fmt.Println("Нет расходов для экспорта")
		return
	}

	fmt.Print("Введите имя файла для экспорта (например export.csv): ")
	scanner.Scan()
	fileName := strings.TrimSpace(scanner.Text())
	if fileName == "" {
		fmt.Println("Имя файла не может быть пустым")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Записываем заголовок CSV
	_, _ = file.WriteString("ID,Description,Amount,Date\n")

	for _, exp := range expenses {
		line := fmt.Sprintf("%d,%s,%.2f,%s\n",
			exp.ID,
			exp.Description,
			exp.Amount,
			exp.Date.Format("2006-01-02"),
		)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка при записи файла:", err)
			return
		}
	}

	fmt.Println("Данные успешно экспортированы в", fileName)
}

func showHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("----------------------------------------")
	fmt.Println("add           - добавить расход")
	fmt.Println("update        - обновить расход по ID")
	fmt.Println("delete        - удалить расход по ID")
	fmt.Println("list          - показать все расходы")
	fmt.Println("summary       - показать общую сумму расходов")
	fmt.Println("month         - показать сумму расходов за месяц")
	fmt.Println("export        - экспортировать данные в CSV")
	fmt.Println("help          - показать список команд")
	fmt.Println("exit / quit   - выйти из программы")
	fmt.Println("----------------------------------------")
}
