package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println(dt)

	var num = make([]int, 5, 8)
	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	showSlice(num)

	var no = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	no = no[0 : len(no)-1]
	fmt.Println(no)
}

func showSlice(x []int) {
	fmt.Printf("len = %d | cap = %d | slice = %v\n", len(x), cap(x), x)
}

func removeSlice(x []int, idx int) []int {
	return append(x[:idx], x[idx+1:]...)
}
