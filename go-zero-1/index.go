package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go is a statically typed")
	fmt.Println(3 * time.Second)
	time.Sleep(3 * time.Second)
}
