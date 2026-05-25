package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в приложение expense-tracker")
	fmt.Println("Приложение может добавлять, удалять и просматривать расходы пользователя. Приложение также может предоставить сводку расходов.")
	fmt.Print("Введите команду: ")

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Сканнер прочитал строку: ", line)
		fmt.Print("Введите команду: ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения: ", err)
	}
}
