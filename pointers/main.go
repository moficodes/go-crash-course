package main

import "fmt"

func main() {
	i := 10
	fmt.Println(&i)
	n := &i
	fmt.Printf("%T\n", n)
	fmt.Println(n)
	fmt.Println(i)
	*n = 15
	fmt.Println(i)

	j := 5
	n = &j
	*n = 6

	fmt.Println(i)

	a := 5
	fmt.Println(a) // value is 5
	pointer(&a)
	fmt.Println(a) // value is 10
	a = 5
	fmt.Println(a) // value is 5
	value(a)
	fmt.Println(a) // value is still 5

	ar := []int{1, 2, 3, 4}
	fmt.Println(ar)
	arrval(ar)
	fmt.Println(ar)
	arrpointer(&ar)
	fmt.Println(ar)

	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3

	fmt.Println(m)
	mapped(m)
	fmt.Println(m)
	m2 := map[int]int{1: 1}
	mapped(m2)
	fmt.Println(m2)
}

func pointer(x *int) {
	*x = 10
}

func value(x int) {
	x = 10
}

func arrval(a []int) {
	a = append(a, 5)
}

func arrpointer(a *[]int) {
	*a = append(*a, 5)
}

func mapped(m map[int]int) {
	m[4] = 4
}
