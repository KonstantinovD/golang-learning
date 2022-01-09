package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"strconv"
)

func main() {

	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is great or equal than b")
	}

	c := 8
	d := 8
	if c < d {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	fmt.Println(".................................")
	fmt.Println()

	e := 8
	f := 50
	g := 78
	funcswitch(e, g)
	funcswitch(f, g)

}

func funcswitch(x, y int) {
	switch x {
	case 9:
		fmt.Println("x = 9")
	case 8:
		{
			fmt.Println("x = 8")
			fmt.Println("x = 8 - twice\n")
		}
	case 7, 6, 5:
		fmt.Println("x = 7 or 6 or 5")
	default:
		fmt.Println("y = " + strconv.Itoa(y))
	}
}
