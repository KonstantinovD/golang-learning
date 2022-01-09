package main

import (
	"fmt"
	"github.com/sqweek/dialog"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	tFile, err := dialog.File().Load()
	if err != nil {
		fmt.Println("Please choose a correct file! ", err)
		os.Exit(1)
	}

	// Исходная версия данных хранится в переменной DATA,
	// которая представляет собой двумерный срез.
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []Entry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Entries = append(Entries, temp)
		}
	}

	// --- Для выполнения всех необходимых инициализаций используется
	// функция template.Must(). Она возвращает данные типа Template
	// - эта структура содержит в себе шаблон, обработанный
	// синтаксическим анализатором
	// --- template.ParseGlob() читает внешний файл шаблона
	t := template.Must(template.ParseGlob(tFile))
	// --- template.Execute() выполняет всю остальную работу, которая
	// включает в себя обработку данных и вывод результатов в нужный
	// файл — в данном случае это os.Stdout
	t.Execute(os.Stdout, Entries)
}
