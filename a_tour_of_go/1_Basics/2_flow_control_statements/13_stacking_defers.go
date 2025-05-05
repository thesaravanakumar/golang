package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ { // prints 9 to 0
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
