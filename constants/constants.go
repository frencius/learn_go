package main

import "fmt"

const s string = "constant string"

func main() {
	fmt.Println(s)

	const n int = 50000000

	const d int = 500000000 / n
	fmt.Println(d)

	const number = 90000
	const result = number / 10
	fmt.Println(int64(result))
}
