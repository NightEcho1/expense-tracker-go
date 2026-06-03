package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	loadExpenses()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в приложение expense-tracker")
	fmt.Println("Приложение может добавлять, удалять и просматривать расходы пользователя. Приложение также может предоставить сводку расходов.")
	fmt.Println("Введите help для просмотра списка команд.")
	fmt.Print("Введите команду: ")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.ToLower(line)
		switch line {
		case "add":
			addExpense(scanner)
		case "update":
			updateExpense(scanner)
		case "delete":
			deleteExpense(scanner)
		case "list":
			listExpenses()
		case "summary":
			showSummary()
		case "summary month", "month":
			showMonthSummary(scanner)
		case "export":
			exportCSV(scanner)
		case "exit", "quit":
			fmt.Println("выход из программы")
			return
		case "help":
			showHelp()
		default:
			fmt.Println("Незивестная команда", line)
		}
		fmt.Print("Введите команду: ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения: ", err)
	}
}
