package main

import (
	// "crypto"
	"fmt"
)

func main() {
	// fmt.Println()
	// enCryption("12345678")
	for i := 1; i < 10; i++ {
		// Add(i)
		// Mul(i)
		for j := 1; j < 10; j++ {
			m := i * j
			fmt.Printf("%d x %d = %d;\n", i, j, m)
		}
		fmt.Print("\n")
	}
}

// func enCryption(text string) {
// 	for i := 0; i < 9; i++ {
// 		fmt.Println(text)
// 	}
// }

// func Add(i int) {
// 	i++
// 	fmt.Println(i)

// }

// func Mul(i int) {
// 	m := i * i
// 	fmt.Println(m)
// 	fmt.Printf("%d x %d = %d;\t", i, i, m)
// 	fmt.Printf("%d x %d = %d;\n", i, i, m)
// }
