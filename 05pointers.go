package main

import "fmt"

func main() {
	// var score int = 34
	// // P will have the reference to the memory location of the score
	// // If score get's updated P will also be updated automatically
	// p := &score
	// if p != nil {
	// 	fmt.Println("P is having a value: ", *p)
	// } else {
	// 	fmt.Println("P is not having a value: ")
	// }

	x := 5
	y := 2
	// square(&x)
	x, y = swap(x, y)
	// swapPtr(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

func swap(x int, y int) (int, int) {
	x, y = y, x
	return x, y
}

func square(x *float64) {
	*x = *x * *x
}

func swapPtr(xPtr *int, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr
}
