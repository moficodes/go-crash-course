package main

import (
	"fmt"
	"math"
)

var User string = "Hello"

type person struct {
	int
}

type aint int

const (
	BigConst      = math.MaxFloat64 * math.MaxFloat64
	data     aint = 10
)

const (
	a0 = iota // 0
	a1 = iota // 1
	a2 = iota // 2
	a3 = iota // 3
)

const (
	b0 = iota // 0
	b1        // 1
	b2        // 2
	b3        // 3
)

const (
	c0 = iota // 0
	c1 = 43
	c2 = 75
	c3 = iota // 3
)

const d0 = iota //0
const d1 = iota //0
const d2 = iota //0
const d3 = iota //0

func main() {
	fmt.Println(a0, a1, a2, a3)
	fmt.Println(b0, b1, b2, b3)
	fmt.Println(c0, c1, c2, c3)
	fmt.Println(d0, d1, d2, d3)

}
