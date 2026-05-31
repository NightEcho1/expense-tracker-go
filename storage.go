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
	//TODO
	//1. Проверить есть файл
	//2.1 Если нету файла, выйти из функции (т.к как записей ещё не было)
	//2.2 Если файл есть, то прочитать его
	//3 Выгрузить данные
}
