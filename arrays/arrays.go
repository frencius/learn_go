package main

import "fmt"

func main() {
	var a [5]int
	b := [5]int{0, 1, 2, 3, 4}
	var twoDimArray [2][3]int

	fmt.Println(a)
	fmt.Println(b)

	a[4] = 10
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimArray[i][j] = i + j
		}
	}

	fmt.Println(twoDimArray)
}
