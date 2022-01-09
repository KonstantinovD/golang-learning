package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var inp string
	fmt.Print("Enter number: ")
	_, err := fmt.Scanf("%s", &inp)

	number, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("Inputted value is not an integer: ", inp)
	} else {
		switch {
		case number < 0:
			fmt.Println("Less than zero!")
		case number > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero!")
		}
	}

	asString := inp
	switch asString {
	case "5":
		fmt.Println("Five!")
	case "0":
		fmt.Println("Zero!")
	default:
		fmt.Println("Do not care!")
	}

	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	// -- NOTE:
	// By default the switch statement matches goes through all the case
	// statement from top to bottom and tries to find the first case
	// expression that matches the switch expression. Once the matching
	//case is found, it exits and does not consider the other cases.
	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point!")
	case email.MatchString(asString):
		fmt.Println("It is an email!")
		fallthrough // enables not to exit switch but to continue
		// going through the cases below
	default:
		fmt.Println("Something else!")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface!")
	default:
		fmt.Println("Not nil interface!")
	}

}
