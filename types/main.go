package main

import (
	"fmt"
	"strconv"
)

type myFloat = float64
type myint int

// func (m myFloat) double() myFloat {

// }

type person struct {
	age  int
	name string
}

type student struct {
	age  int // try swapping these fields
	name string
}

func (m myint) double() myint {
	return m * 2
}

func main() {
	var a uint8 = 1
	var b uint16 = 2
	var c uint32 = 3
	var d uint64 = 4

	var e int8 = 1
	var f int16 = 2
	var g int32 = 3
	var h int64 = 4

	var i float32 = 1.1
	var j float64 = 1.2

	var k bool = true

	var l byte = 'c'
	var m byte = 99
	var n rune = '😀'
	var o rune = 128512

	var p string = "hello"

	var arr = [5]int{}
	var slice = []int{}
	var maps = map[int]int{}

	type empty struct{}
	var x = empty{}

	var y interface{} = x

	var function = func() int {
		return 42
	}

	var aliasInt myint = 1
	var aliasFloat myFloat = 5.0

	fmt.Printf("Type of a=%T, b=%T, c=%T, d=%T\n", a, b, c, d)
	fmt.Printf("Type of e=%T, f=%T, g=%T, h=%T\n", e, f, g, h)
	fmt.Printf("Type of i=%T, j=%T\n", i, j)
	fmt.Printf("Type of k=%T\n", k)
	fmt.Printf("Type of l=%T, m=%T, n=%T, o=%T\n", l, m, n, o)
	fmt.Printf("Type of p=%T\n", p)
	fmt.Printf("Type of arr=%T\n", arr)
	fmt.Printf("Type of arr=%T\n", slice)
	fmt.Printf("Type of arr=%T\n", maps)
	fmt.Printf("Type of x=%T\n", x)
	fmt.Printf("Type of y=%T\n", y)
	fmt.Printf("Type of function=%T\n", function)
	fmt.Printf("Type of aliasInt=%T, aliasFloat=%T\n", aliasInt, aliasFloat)
	conversion()

	definedTypeConversion(aliasInt)
	definedTypeConversion(5)

	takesInt(5)
	// takesInt(aliasInt) // this does not work

	takesFloat(5)
	takeInt64(5)
	takeInt8(5) // try passing > 128
}

func conversion() {
	var i int = 355 // type int
	f := float64(i) // type float64
	b := i != 0     // type bool
	s := strconv.Itoa(i)
	by := byte(i)

	ui8 := 250
	i8 := int8(ui8)

	fmt.Printf("%d\t%f\t%t\t%s\t%c\t%d\t%d\n", i, f, b, s, by, ui8, i8)

	p := person{age: 20, name: "john"}
	std := student{age: 15, name: "paul"}

	p2 := person(std)
	s2 := student(p)

	// p2 := *(*person)(unsafe.Pointer(&std))
	// s2 := *(*student)(unsafe.Pointer(&p))

	fmt.Printf("Type of p=%T, std=%T, p2=%T, s2=%T\n", p, std, p2, s2)
}

func definedTypeConversion(mi myint) {
	fmt.Printf("Type of mi=%T\n", mi)
}

func takesInt(i int) {
	fmt.Printf("Type of i=%T\n", i)
}

func takesFloat(f float32) {
	fmt.Printf("Type of i=%T\n", f)
}

func takeInt8(i int8) {
	fmt.Printf("Type of i=%T\n", i)
}

func takeInt64(i int64) {
	fmt.Printf("Type of i=%T\n", i)
}
