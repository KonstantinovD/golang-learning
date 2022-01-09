package main // определение пакета для текущего файла

import ( // подключение пакета fmt
	"fmt"
)

func identation() {
	fmt.Println()
	fmt.Println("..................")
	fmt.Println()
}

func main() {
	// Фактически в Go есть только один цикл - цикл for, который может принимать разные формы
	/* for [инициализация счетчика]; [условие]; [изменение счетчика]{
		// действия
	} */
	for i := 1; i < 10; i++ {
		fmt.Print(i * i)
		fmt.Print(" ")
	}
	identation()

	// можно вынести объявление переменной вовне:
	var j = 1
	for ; j < 10; j++ {
		res := j + 3
		if res >= 8 && res < 11 {
			continue
		}
		fmt.Print(res)
		fmt.Print(" ")
	}
	identation()

	// убрать изменение счетчика в тело цикла и оставить только условие
	// похоже на цикл while
	// Если цикл использует только условие, то можно убрать обе ";"
	var k = 1
	for k < 10 {
		fmt.Print(k * 2)
		fmt.Print(" ")
		k++
		if k*2 > 13 {
			break
		}
	}
	identation()

	// Циклы могут быть вложенными:
	for iter1 := 1; iter1 < 10; iter1++ {
		for iter2 := 1; iter2 < 10; iter2++ {
			fmt.Print(iter1*iter2, "\t")
		}
		fmt.Println()
	}
	identation()

	//Перебор массивов
	/* for индекс, значение := range массив{
		// действия
	} */
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	identation()

	var phrase string = "Если мы не планируем использовать значения" +
		"или индексы элементов, то мы можем вместо них указать прочерк"
	fmt.Println(phrase)
	for _, value := range users {
		fmt.Print(value)
		fmt.Print(" ")
	}
	identation()

	// длинну массива можно получить функцией len()
	fmt.Print("длина массива \"friends\": ")
	fmt.Println(len(users))
}
