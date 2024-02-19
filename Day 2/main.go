package main

import (
	"fmt"
	"log"
)

func main() {

	f := func(int, int) int {
		return 1 + 1
	}

	a := func(x, y int) int {
		return x * y
	}

	fmt.Println(f(2, 5))
	f = a
	fmt.Println(f(2, 5))

}

func okruglenie() {

	func() int {
		a++
		log.Println(a)
		return a
	}()

}
