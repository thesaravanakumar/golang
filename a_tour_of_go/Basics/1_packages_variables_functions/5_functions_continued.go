package main

import "fmt"

func add(x, y int) int { // function parameters share a type, you can omit the type from all but the last
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
