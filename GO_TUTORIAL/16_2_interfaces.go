package main

import (
	"fmt"
)

/*
type имя_интерфейса interface{
    определения_функций
}
*/
// --- An interface is two things:
// 1) it is a set of methods
// 2) but it is also a type

type vehicle interface {
	move() // Данный метод представляет действие или поведение,
	// которые могут реализовывать другие объекты.
}

// интерфейсы определяют только поведение, но не фактические
// реализации. Это уже работа объекта, реализующего данный интерфейс.
// Интерфейс - это абстракция, а не конкретный тип, нельзя напрямую
// создать объект интерфейса
/* var v vehicle = vehicle{} */ // ! ОШИБКА

// Чтобы тип данных соответствовал некоторому интерфейсу, данный тип
// должен реализовать в виде методов ВСЕ функции этого интерфейса
type armoredVehicle interface {
	move()
	fire()
}

/* func (av armoredVehicle) fire(){ } */
// unable to make default interface methods

type Car struct{}

type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}
func (a Aircraft) fire() {

}

func main() {
	var vehicles []vehicle = []vehicle{}
	var plane1 Aircraft = Aircraft{}
	var plane2 Aircraft = Aircraft{}
	var car1 Car = Car{}
	vehicles = append(vehicles, plane1, car1, plane2)
	for _, vh := range vehicles {
		vh.move()
	}

	// --- обязательно надо реализовывать все методы интерфейса
	var armored []armoredVehicle = []armoredVehicle{}
	armored = append(armored, plane1, plane2)
	/* armored = append(armored, car1) */ // ! Ошибка
	// ERROR:
	// Car does not implement armoredVehicle (missing fire method)
	for _, arm := range armored {
		arm.fire()
	}
	fmt.Println("------------")
	fmt.Println()
	// Интерфейсы также могут быть использованы в качестве полей

	// Вложенные интерфейсы - Одни интерфейсы могут содержать другие:
	// В этом случае для соответствия подобному интерфейсу типы данных
	// должны реализовать все его вложенные интерфейсы
	var myFile ReaderWriter = &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
	writeToStream(myFile, "lolly bomb")
	readFromStream(myFile)
}

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}
func readFromStream(reader Reader) {
	reader.read()
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}
