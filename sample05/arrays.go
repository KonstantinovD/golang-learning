package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
	"strconv"
)

func main() {

	// Массив определяется следующим способом:
	// var numbers [число_элементов]тип_элементов
	var numbers [5]int //элементы массива инициализ. значениями по умолчанию
	fmt.Println(numbers[4])

	var numbers2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2) // [1 2 3 4 5]
	var numbers3 [5]int = [5]int{1, 2}
	fmt.Println(numbers3)          // [1 2 0 0 0]
	numbers4 := [5]int{1, 2, 3, 4} //сокращенное определение переменной массива
	fmt.Println(numbers4)          // [1 2 3 4 0]

	// Если в квадратных скобках вместо длины указано троеточие,
	// то длина массива определяется, исходя из количества
	// переданных ему элементов:
	var numbers5 = [...]int{1, 2, 3, 4, 5} // длина массива 5
	numbers6 := [...]int{1, 2, 3}          // длина массива 3
	fmt.Println(numbers5)                  // [1 2 3 4 5]
	fmt.Println(numbers6)                  // [1 2 3]

	// При этом длина массива является частью его типа.
	// Cледующие два массива представляют разные типы данных,
	// хотя они и хранят данные одного типа:
	var numbers7 [3]int = [3]int{1, 2, 3}
	var numbers8 [4]int = [4]int{1, 2, 3, 4}
	/* numbers7 = numbers8 */ // ! Ошибка
	var ignore string = strconv.Itoa(numbers7[0]) + strconv.Itoa(numbers8[0])
	fmt.Println(ignore)

	// Индексы в массиве фактически выступают в качестве ключей для значений
	// В прицнипе мы можем явным образом указать, значения для ключей
	// При этом числовые ключи необязательно располагать в порядке возрастания
	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue

	fmt.Println("\nmatrix:")
	var matrix [4][3]int = [4][3]int{{1, 2, 3}, {4, 5}, {7, 8, 9}, {10, 11, 12}}

	fmt.Println(matrix[2][2]) // = 9
	fmt.Println(matrix[1][2]) // = 0
	printMatrix(matrix)

	// dynamic arrays - append function
	friends := []string{"Adam"}
	friends = append(friends, "Rahul") // Add one friend or one string
	for i := 0; i < 2; i++ {
		fmt.Print(friends[i] + " ")
	}
	fmt.Println()
	friends = append(friends, "Angelica", "Rashi") // Add multiple friends or multiple strings
	for i := 0; i < 4; i++ {
		fmt.Print(friends[i] + " ")
	}
	fmt.Println("\nlenght of friends: ", len(friends))
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
}

func printMatrix(matrix [4][3]int) {
	fmt.Println("pinting all the matix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
