package interfaces

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	sidelength float64
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	Hesap() float64
}

func (c circle) Hesap() float64 {
	return math.Pi * c.radius
}

func (r rectangle) Hesap() float64 {
	return r.height * r.width * 2
}

func (s square) Hesap() float64 {
	return s.sidelength * 4
}

func school(s shape) {
	fmt.Println("Measurement of area :", s.Hesap())
}
func Demo2() {
	circle := circle{radius: 12}
	square := square{sidelength: 12}
	rectangle := rectangle{width: 5, height: 3}
	school(circle)
	school(square)
	school(rectangle)
}
