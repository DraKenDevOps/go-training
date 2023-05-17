package main

import "fmt"

func main() {
	resA := Add(2, 2)
	resS := Sub(4, 2)
	resM := Mul(9, 8)
	resD := Divd(81, 9)
	// resM := Modd(86, 5)

	fmt.Println(resA, resS, resM, resD)

	// custom while loop
	x := 0
	for x < 5 {
		x++
		fmt.Printf("While loop index: %d;\n", x)
	}

	stack := []string{"MERN", "MEVN"}

	for idx, item := range stack {
		fmt.Printf("%d: %s\t", idx+1, item)
	}
}

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Divd(a, b int) int {
	return a / b
}

// func Modd(a, b int) int {
// 	r := a + b
// 	return r
// }
