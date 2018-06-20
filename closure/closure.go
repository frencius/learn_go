package main

import "fmt"

func main() {
	newInt := intSeq()

	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	newInts := intSeq()

	fmt.Println(newInt())
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
