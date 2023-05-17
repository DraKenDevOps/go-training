package main

import (
	"fmt"
)

func main() {
	msg := "Spiderman 2099"
	var msg_pointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msg_pointer)
	changeMessage(&msg)
	fmt.Println("Message changed: ")
	fmt.Println(msg)
	changeMessage1(&msg, "Buddha 7 steps")
	fmt.Println("Message changed 1: ")
	fmt.Println(msg)
	fmt.Println("Pointer *msg_pointer: ")
	fmt.Println(*msg_pointer)
}

func changeMessage(m_pointer *string) {
	*m_pointer = "Draken DevOps Engineer"

}

func changeMessage1(m_pointer *string, new_msg string) {
	*m_pointer = new_msg
}
