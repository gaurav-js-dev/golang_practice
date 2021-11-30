package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sidelength float64
}

func main() {
	sq := square{sidelength: 5}
	tri := triangle{height: 10, base: 20}
	printArea(sq)
	printArea(tri)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}
