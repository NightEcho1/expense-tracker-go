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
			fmt.Println("Добавить расход: ")
		case "update":
			fmt.Println("Обновить запись об расходе: ")
		case "delete":
			fmt.Println("Удалить запись об расходе: ")
		case "viewall":
			fmt.Println("Посмотреть все расходы: ")
		case "allsummary":
			fmt.Println("Посмотреть информацию обо всех расходов: ")
		case "monthsummary":
			fmt.Println("Посмотреть информацию обо всех расходов за месяц: ")
		case "export":
			fmt.Println("Экспортировать в файл .csv")
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
