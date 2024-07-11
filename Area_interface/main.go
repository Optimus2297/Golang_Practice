package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {

	s := square{sideLength: 5.7}
	t := triangle{base: 4.5, height: 7.8}

	printArea(s)
	printArea(t)

}

func printArea(sh shape) {
	fmt.Println("Area of the shape is :", sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}