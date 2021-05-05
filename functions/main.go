package main

import "fmt"

func init() {
	fmt.Println("hello from init 1")
}

func init() {
	fmt.Println("hello from init 2")
}

func init() {
	fmt.Println("hello from init 3")
}

func main() {
	fmt.Println("Hello World")
	q, r := NamedDivide(5, 3)
	fmt.Printf("Divide 5 / 3, Quotient: %d Remainder: %d\n", q, r)
	fmt.Printf("Bill: %.2f\n", Bill(8.75, 10.2, 3.7))
	Greet()

	Greet := func() string {
		return "shadow outer Greet"
	}

	TakeFunc(GiveFunc())
	TakeFunc(Greet)
}

func Greet() {
	fmt.Println("hello")
}

// this will fail
// func Greet(m string) {
// 	fmt.Println(m)
// }

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, int) {
	return a / b, a % b
}

func NamedDivide(a, b int) (quotient int, remainder int) {
	quotient, remainder = a/b, a%b
	return
}

func TakeFunc(take func() string) {
	data := take()
	fmt.Println(data)
}

func GiveFunc() func() string {
	f := func() string {
		return "give"
	}
	return f
}

func Bill(tax float64, items ...float64) (total float64) {
	for _, item := range items {
		total += item
	}

	total *= (1 + tax*.01)
	return
}
