package main

import "fmt"

func doublesquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	var y int
	fmt.Print("Enter value: ")
	_, err := fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))

	double := func(s int) int {
		return s * 2
	}
	fmt.Println("The double of", y, "is", double(y))

	fmt.Println(doublesquare(y))
	d, s := doublesquare(y)
	fmt.Println(d, s)
}
