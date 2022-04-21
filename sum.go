package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func sumX(a, b int) int {
	return a + b + a
}
