package main

import "fmt"

func main() {
	var numbers [3]string
	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"
	fmt.Println(numbers)

	var colors = [3]string{"one", "two", "three"}
	fmt.Println(colors)
	fmt.Println(colors[1])
	fmt.Println(len(colors))

}
