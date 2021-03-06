package main

import (
	"fmt"
	b "leaning/BOOK/285_go_packages/02_init_2"
	a "leaning/BOOK/285_go_packages/02_init_example"
)

// каждом Go-пакете может присутствовать закрытая функция с именем
// init(), которая автоматически выполняется в начале выполнения пакета

// Функция init() — это закрытая функция, то есть ее нельзя вызвать
// извне пакета, к которому она принадлежит. Кроме того, поскольку
// пользователь пакета не имеет контроля над функцией init(), вам
// следует хорошо подумать, прежде чем использовать ее в общедоступных
// пакетах или изменять в ней любое глобальное состояние.

func init() {
	fmt.Println("init() manyInit")
}

// (see output below)
func main() {
	a.FromA()
	b.FromB()
} //       |
//       V

// init() a
// init() b
// init() manyInit
// fromA()
// fromB()
// fromA()

// --- функция init() для пакета a выполняется только один раз, несмотря
// на то что пакет импортируется дважды, двумя разными пакетами.
// Так как сначала выполняется блок import из manyInit.go, функции
// init() пакетов a и b запускаются раньше, чем функция init() файла
// manyInit.go, что вполне оправданно.
