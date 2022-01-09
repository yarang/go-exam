package main

import "fmt"

func execute() func(x int) int {
	var a = 3
	var b = 3
	return func(x int) int {
		fmt.Printf("a: %d, ", a)
		fmt.Printf("b: %d, ", b)
		fmt.Printf("x: %d\n", x)
		return a*b + x
	}
}

func main() {
	var f = execute()

	fmt.Printf("f(10): %d\n", f(10))
	fmt.Printf("f(11): %d\n", f(11))
}
