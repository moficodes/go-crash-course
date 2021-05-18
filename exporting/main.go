package main

import (
	"fmt"

	"github.com/moficodes/go-crash-course/exporting/print"
)

func main() {
	m := print.Message{
		Greeting: "Hello",
	}
	fmt.Println(print.Number)
	fmt.Println(m.voice)
	// fmt.Println(print.data)
}

func Export() int {
	return 42
}
