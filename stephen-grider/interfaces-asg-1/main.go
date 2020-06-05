package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func main() {
	tri := triangle{5, 7}
	sqr := square{6}

	fmt.Println(tri.getArea(), sqr.getArea())
}
