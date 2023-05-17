package main

import "fmt"

var name string = "Draken"
var count int = 0

func main() {
	fmt.Println("Call inCrement()")
	inCreament()
}

func inCreament() {
	fmt.Println("Run inCrement()")
	fmt.Print(name)
	count++
	fmt.Print(count)
	count++
	fmt.Print(count)
	count++
	fmt.Print(count)
	count++
	fmt.Print(count)
	count++
	fmt.Print(count)
}
