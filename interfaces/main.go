package main

import (
	"fmt"
	"math"
	"net/http"
)

type bob int

func (b bob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {
	circle := Circle{5.0}
	r1 := Rectangle{4.0, 5.0}
	r2 := &Rectangle{10.0, 15.0}
	area := circle.Area()
	fmt.Println(area)
	GetArea(&circle)
	GetArea(r1)
	GetArea(r2)

	var x interface{} = 10
	fmt.Println(x)

	str, ok := x.(string)
	if !ok {
		fmt.Println("could not convert")
	}

	fmt.Printf("value: %v, Type: %T\n", x, x)
	x = 1
	fmt.Printf("value: %v, Type: %T\n", x, x)

	fmt.Printf("value: %v, Type: %T\n", str, str)

	var y interface{} = "go"
	switch v := y.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
