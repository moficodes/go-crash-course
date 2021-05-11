package main

import (
	"fmt"
)

func main() {
	adder := func(a int, b int) int {
		return a + b
	}

	fmt.Println(adder(5, 6))

	counter1 := counter()
	counter2 := counter()

	fmt.Println(counter1())
	fmt.Println(counter2())

	deferredPrint()
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func deferredPrint() {
	defer fmt.Println("defer 1")
	fmt.Println("regular thing")
	defer fmt.Println("defer 2")
}
