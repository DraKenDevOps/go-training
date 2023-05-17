package main

import (
	"fmt"
)

func main() {
	var heros = []string{"THOR", "CAP", "IRONMAN", "SPIDERMAN", "SPIDERMAN2099"}
	var users [2]string
	fmt.Println(heros)
	for i := 0; i < len(heros); i++ {
		fmt.Println(heros[i])
	}

	saySorry(20)
	fmt.Println("END")

	var num []int = []int{1, 2, 3, 4}
	fmt.Println(num)

	for i, item := range num {
		fmt.Println("Loop of num")
		fmt.Println(item + i)
	}

	users[0], users[1] = "Owen", "Kin"

	fmt.Println(users)
}

func saySorry(len int) {
	for i := 0; i < len; i++ {
		fmt.Println("Sorry 🥺.")
	}
}
