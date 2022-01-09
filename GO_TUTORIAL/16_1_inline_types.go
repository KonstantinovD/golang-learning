package main

import (
	"fmt"
)

// --- Встраиваемые типы
//Обычно, поля struct представляют отношения принадлежности (включения).
// Например, у Circle есть radius.
// Предположим, у нас есть структура Person:

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// И если хотим создать новую структуру Android (v1), то сделаем так:

type AndroidFirstVersion struct {
	Person Person
	Model  string
}

// Это будет работать, но мы можем захотеть создать другое отношение.
// Сейчас у андроида «есть» личность, можем ли мы описать отношение
// андроид «является» личностью? Go поддерживает подобные отношения
// с помощью встраиваемых типов, также называемых анонимными полями.
// Выглядят они так:

type Android struct {
	Person
	Model string
}

func main() {
	// --- Анонимная структура доступна через имя типа:
	a := new(Android)
	a.Person.Name = "Elisa"
	a.Person.Talk()
	// Но мы также можем вызвать метод/поле Person прямо из Android:
	a.Name = "Martin"
	a.Talk()

	// --- Конструктор для встраиваемого типа
	// !!! Note !!!
	// cannot use 	b := Android{Model: "Model 2", Name: "Martin"}
	// but can use:
	c := Android{Model: "Model 3", Person: Person{Name: "Kevin"}}
	c.Talk()

}
