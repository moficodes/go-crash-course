package main

import "fmt"

var a = 10
var b int8 = 122
var c, d = 100, 200

const PI = 100

func main() {
	var a = 10
	var b = 1.5
	var c byte = 'c'
	var d = true
	var e = '😀'
	var f = "hello"
	var g = 1 + 3i

	var _1åœπbool = 10

	fmt.Printf("%T\t%T\t%T\t%T\t%T\t%T\t%T\t%T\n", a, b, c, d, e, f, g, _1åœπbool)

}
