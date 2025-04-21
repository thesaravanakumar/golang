package main

import "fmt"

func add(x int, y int) int { // type comes after variable name
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
