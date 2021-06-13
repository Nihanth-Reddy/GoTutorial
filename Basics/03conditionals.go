package main

import "fmt"

func main() {

	var isLoggedIn bool = true
	var balance int = 10

	if isLoggedIn || balance > 15 {
		fmt.Println("Show cart page")
		var len, error = fmt.Println("Show cart page")
		if error == nil {
			fmt.Printf("length is %v, %T", len, len)
		}
	} else {
		fmt.Println("Kindly please log in")
	}

}
