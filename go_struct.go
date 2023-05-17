package main

import "fmt"

func main() {
	var p1 product

	p1.id = 1
	p1.title = "Ethereum"
	p1.price = 5000000

	show(p1)

	p1.id = 2
	p1.title = "Bitcoin"
	p1.price = 8000000

	update(&p1)
	show(p1)

	// p1 = p1.clear()
	p1 = p1.priceFall(20)
	show(p1)
}

type product struct {
	id    int
	title string
	price int
}

// type Product struct {
// 	Id int
// 	Title string
// 	Price int
// }

func show(p product) {
	fmt.Println(p)
}

func update(p *product) {}

func (p product) clear() product {
	p.price = 0
	p.title = ""
	return p
}

// Percent
func (p product) priceFall(percent int) product {
	var fall int = p.price * percent / 100
	fmt.Println("Cut: ", fall)
	p.price = p.price - fall
	return p
}
