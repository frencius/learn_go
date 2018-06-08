package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write i as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	i = 20
	fmt.Print("Write i as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("another number")
	}

	timer := time.Now().Weekday()
	switch string(timer) {
	case "Saturday", "Sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("Weekday")
	}

	switch time.Now().Weekday() {
	case time.Friday:
		fmt.Println("TGIF")
	default:
		fmt.Println("TGINF")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	case t.Hour() > 12 && t.Hour() < 18:
		fmt.Println("noon")
	default:
		fmt.Println("evening")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it's boolean")
		case int:
			fmt.Println("it's integer")
		case float64:
			fmt.Println("it's float")
		default:
			fmt.Printf("it's %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(123)
	whatAmI(12.2)
	whatAmI("jajanan")
	whatAmI('1')

	var rere complex64
	whatAmI(rere)
	fmt.Println(rere)
}
