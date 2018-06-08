package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j < 11; j++ {
		fmt.Print(j)
	}

	fmt.Println()

	ii := 13
	for {
		fmt.Println(ii)
		break
	}

	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
