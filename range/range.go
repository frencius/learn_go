package main

import "fmt"

// add comment range
func main() {
	nums := []int{2, 3, 4}
	fmt.Println("nums:", nums)

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"satu": "1", "dua": "2"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("%s\n", k)
	}

	for _, v := range kvs {
		fmt.Printf("%s\n", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
