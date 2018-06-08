package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)

	sc := make([]string, 10)
	fmt.Println(sc)

	s[0] = "ini "
	s[1] = "adalah "
	s[2] = "go"

	fmt.Println("set:", s)
	fmt.Println("get:", s[0])
	fmt.Println("len:", len(s))

	s = append(s, "lang")
	s = append(s, "uage", "by Google")

	fmt.Println("set:", s)
	fmt.Println("get:", s[0])
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("s:", s)
	fmt.Println("c:", c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	ss := [5]string{"g", "h", "i"}
	fmt.Println(ss)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var becak [1]int
	becak[0] = 1
	fmt.Println(becak)

	fmt.Println("======slices=====")

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		//twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	var twoDi [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoDi[i][j] = i + j
		}
	}
	fmt.Println(twoDi)

}
