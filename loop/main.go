package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; ; {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for ; ; i++ {
		if i >= 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}

	i = 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; ; i++ {
		if i >= 5 {
			break
		}
		fmt.Println(i)
	}

	i = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i, j := 0, 10; i != j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}
}
