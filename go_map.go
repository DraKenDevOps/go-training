package main

import (
	"fmt"
)

func main() {
	var num = map[string]int{"One": 1, "Two": 2, "Three": 3}
	fmt.Println(num)

	var heros = map[string]string{"id": "ASDFGH", "name": "Kin", "age": "22"}
	fmt.Println(heros)

	var stacks = make(map[string]map[string]int)

	stacks["react"] = map[string]int{"price": 500}
	stacks["vue"] = map[string]int{"price": 500}
	stacks["angular"] = map[string]int{"price": 500}
	stacks["docker"] = map[string]int{"price": 500}
	stacks["Kubernetes"] = map[string]int{"price": 500}
	// stacks["react"]["price"] = 500
	// stacks["vue"]["price"] = 500
	// stacks["angular"]["price"] = 500
	// stacks["docker"]["price"] = 500
	// stacks["Kubernetes"]["price"] = 500

	fmt.Println(stacks)
}
