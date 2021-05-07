package main

import "fmt"

func main() {
	x := 5*10 + 6/3 | 7%4&15 // same as (5*10) + (6/3) | (7%4)&15 = 50 + 2 | 3&15 = 52 | 3 =
	fmt.Println(x)

	a := true && false || true && !false
	fmt.Println(a)
}
