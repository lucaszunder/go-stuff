package main

import "math"

type Square struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.width * s.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.width * (t.height / 2)
}
