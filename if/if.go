package main

import "fmt"

func main() {

	if 10%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if number := 9; number < 0 {
		fmt.Println("Number is negative")
	} else if number < 10 {
		fmt.Println("number is 1 digit")
	} else {
		fmt.Println("number is more than 1 digit")
	}

}
