package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	s := []int{1, 2, 3, 4}
	sum(s...)
}

func sum(sums ...int) {
	fmt.Print(sums, " ")
	total := 0
	for _, num := range sums {
		total += num
	}
	fmt.Println(total)
}
