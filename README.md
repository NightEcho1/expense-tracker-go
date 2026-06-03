expense-tracker
Описание / Description

Русский:
Expense Tracker — это простое консольное приложение для учёта личных расходов. С его помощью можно:

Добавлять расходы с описанием и суммой
Изменять или удалять записи
Просматривать все расходы
Суммировать расходы за месяц или за весь период
Экспортировать данные в CSV для анализа

Все данные сохраняются в файле expenses.json, так что при следующем запуске информация не теряется.

English:
Expense Tracker is a simple console application for tracking personal expenses. With it, you can:

Add expenses with description and amount
Update or delete records
View all expenses
Summarize expenses by month or overall
Export data to CSV for further analysis

All data is saved in the expenses.json file, so your information is kept between sessions.
Установка / Installation

Русский:

Скачайте или клонируйте репозиторий
В терминале (PowerShell, CMD) перейдите в папку проекта
Соберите программу:
go build -o expense-tracker.exe
Запустите:
.\expense-tracker.exe

English:

Download or clone the repository
Open terminal (PowerShell, CMD) in the project folder
Build the program:
go build -o expense-tracker.exe
Run the program:
.\expense-tracker.exe
Команды / Commands

Русский:

Команда	Что делает
add	Добавить новый расход
update	Изменить существующий расход по ID
delete	Удалить расход по ID
list	Показать все расходы
summary	Показать общую сумму расходов
month или summary month	Показать расходы за месяц
export	Экспортировать данные в CSV
help	Показать список команд
exit / quit	Выйти из программы

English:

Command	What it does
add	Add a new expense
update	Update an existing expense by ID
delete	Delete an expense by ID
list	Show all expenses
summary	Show total expenses
month or summary month	Show monthly expenses
export	Export data to CSV
help	Show the list of commands
exit / quit	Exit the program
Формат данных / Data format

Все расходы сохраняются в JSON (expenses.json):

[
  {
    "id": 1,
    "description": "Coffee",
    "amount": 3.50,
    "date": "2026-06-03T17:30:00"
  }
]
Экспорт в CSV / Export to CSV

При экспорте создаётся CSV файл с такими столбцами:

ID,Description,Amount,Date
1,Coffee,3.50,2026-06-03