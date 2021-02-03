package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() (area float64)
}

func main() {
	tr := triangle{height: 10.0, base: 5.0}
	sq := square{sideLength: 10.0}

	printArea(tr)
	printArea(sq)
}

func (tr triangle) getArea() (area float64) {
	return (tr.height * tr.base) / 2
}

func (sq square) getArea() (area float64) {
	return sq.sideLength * 2
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
