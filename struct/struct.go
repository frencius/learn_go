package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 12})
	fmt.Println(person{name: "bobay", age: 12})
	fmt.Println(person{name: "bobayo"})
	fmt.Println(person{age: 12})
	fmt.Println(&person{name: "ann", age: 120})

	s := person{name: "leo", age: 12}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 1122
	fmt.Println(sp.age)

}
