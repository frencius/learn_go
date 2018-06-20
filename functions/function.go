package main

import "fmt"

func main() {
	fmt.Println(plus(10, 2))
	fmt.Println(plusPlus(10, 2, plus(10, 2)))
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
