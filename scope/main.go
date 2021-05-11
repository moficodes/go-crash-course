package main

import "fmt"

type person struct {
	name string
}

var p = person{name: "R Doe"}

func main() {
	p1 := person{name: "Jane Doe"} // this is global person

	type person struct {
		name string
		age  int
	}

	p2 := person{ // this is shadowed person
		name: "John Doe",
		age:  27,
	}

	fmt.Printf("Type of p: %+v\tp1: %v\tp2: %v\n", p, p1, p2)
	blocks()
	scopes()
	shadowing()
}

func blocks() {
	i := 10
	{
		i := 5
		fmt.Println(i) // i is 5
	}
	fmt.Println(i) // i is 10
}

var y = 100

func scopes() {
	x := 10
	var z int
	{
		fmt.Println(x)
		y := 15
		fmt.Println(y)
		z = 20
	}
	fmt.Println(z)
	fmt.Println(y)

}

func shadowing() {
	x := 10
	{
		x := 15
		{
			x := 20
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)

}
