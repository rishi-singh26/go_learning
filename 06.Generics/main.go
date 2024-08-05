package main

import "fmt"

func main() {
	result := add("a", "b")
	fmt.Println(result)
}

func add[T int | string | float64](a, b T) T {
	return a + b
}
