package main

import "fmt"

func main() {
	var t1 string = "test"
	var t2 int = 10
	var t3 bool = true

	// implicit declaration
	t4 := "Owen"
	t1 = "codemobiles"
	const t5 string = "*ornhub"
	// t5 = "yyy"
	name := "Draken"

	fmt.Println(t1, t2, t3, t4, t5, name)
	// fn1("Owen")
	// fn2("hello", 11, false)

	// result := fn3("hello", 11, false)
	// fmt.Println(result)

	// // Array
	// var a1 []int = []int{1, 2, 3, 4}
	// fmt.Println(a1)

	// a2 := []string{"react", "vue", "flutter"}
	// fmt.Println(a2)
	// fmt.Println(len(a2))
	// fmt.Println(a2[1])

	// // Slice
	// var s1 = make([]int, 3)
	// s1[0] = 8
	// s1[1] = 7
	// s1[2] = 6
	// s1 = append(s1, 5)
	// s1 = append(s1, 4)
	// s1 = append(s1, 3)
	// s1 = s1[0 : len(s1)-1]
	// fmt.Println(s1)

	// // Map
	// var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println("", numbers["two"])

	// courses := map[string]string{"v": "vue", "r": "react"}
	// if courses["v"] == "vue" {
	// 	fmt.Println("This is Vue")
	// } else {
	// 	fmt.Println("This NOT Vue")
	// }

	// if result := 11; result == 10 {
	// 	fmt.Println("OK")
	// }
}

// func fn1(p1 string) {
// 	fmt.Printf("Hey %s\n", p1)
// }

// func fn2(p1 string, p2 int, p3 bool) {
// 	fmt.Printf("Hey %s, %d, %v\n", p1, p2, p3)
// }

// func fn3(p1 string, p2 int, p3 bool) string {
// 	fmt.Printf("Hey %s, %d, %v\n", p1, p2, p3)
// 	return "I am return Owen"
// }
