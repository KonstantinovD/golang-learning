package init_b

import (
	"fmt"
	a "leaning/BOOK/285_go_packages/02_init_example"
)

// пакет b должен импортировать пакет a, поскольку в нем применяется
// функция a.FromA(). И в a, и в b есть функция init().
func init() {
	fmt.Println("init() b")
}
func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}
