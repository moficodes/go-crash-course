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
}
