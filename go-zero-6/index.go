package main

import (
	"fmt"
)

func main() {
	val := 10

	if val == 10 {
		fmt.Print(val)
	} else {
		fmt.Print("val < 10 || val > 10")
	}
}
