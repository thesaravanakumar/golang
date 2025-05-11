package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ { // prints 9 to 0
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0