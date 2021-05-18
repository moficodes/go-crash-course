package main

import (
	"fmt"
	"math"
)

type Wheel struct {
	Circle
	Material string
	Color    string
	X        float64
}

func (w Wheel) Info() {
	fmt.Printf("Wheel made of %s and has %s color\n", w.Material, w.Color)
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

type Shape interface {
	Area() float64
	Area2() float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Area2() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Area(c Circle) float64 {
	return c.Radius * c.Radius * math.Pi
}

func ShapeArea(s Shape) float64 {
	return s.Area()
}

func main() {
	c1 := Circle{
		X:      15.0,
		Y:      12.0,
		Radius: 8.5,
	}

	c2 := Circle{15.0, 12, 8.5}

	fmt.Println(c1 == c2)

	fmt.Println(c1.Area() == c1.Area2())

	w1 := Wheel{
		Circle: Circle{
			Radius: 10.0,
			X:      15.0,
		},
		Material: "Rubber",
		Color:    "Black",
		X:        5.0,
	}
	w1.Info()

	fmt.Println(w1.Area())
	fmt.Println(w1.X)
	fmt.Println(w1.Circle.X)
}
