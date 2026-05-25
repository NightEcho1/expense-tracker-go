package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в приложение expense-tracker")
	fmt.Println("Приложение может добавлять, удалять и просматривать расходы пользователя. Приложение также может предоставить сводку расходов.")
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
			deleteEpxense(scanner)
		case "list":
			listExpenses()
		case "summary":
			showMonthSummary(scanner)
			//TODO
			// 1. Сделать цикл в котором будет пробег по всем месяцам, что в результате будет давать год
			// 2. Подумать над тем как сделать вывод (Каждый месяц отдельно, или просто всё сразу)
		case "summary month", "month":
			showMonthSummary(scanner)
		case "export":
			exportCSV()
		case "exit", "quit":
			fmt.Println("выход из программы")
			return
		default:
			fmt.Println("Незивестная команда", line)
		}
		fmt.Print("Введите команду: ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения: ", err)
	}
}
