package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x = 1
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// for i := 1; i < 100; i++ {
	// 	fizzbuzz(i)
	// }
	fmt.Println(minRand(5000000000000000000))
}

func fizzbuzz(n int) {
	if n%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if (n % 5) == 0 {
		fmt.Println("Buzz")
	} else if (n % 3) == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(n)
	}
}

func minRand(min int) int {
	rand.Seed(time.Now().UnixNano())
	if v := rand.Int(); v > min {
		return v
	}
	return min
}

func isEven(n int) bool {
	if n%2 == 1 {
		return false
	}
	return true
}
