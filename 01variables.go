package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"

	fmt.Println(superman)

	thor := "I am son of ODIN"
	fmt.Println(thor)

	thorRating := 3
	fmt.Printf("%v, %T\n", thorRating, thorRating)

	var Ironman, CapAmerica string = "I am Tony Stark", "I am Steve Rogers"
	fmt.Println(Ironman, CapAmerica)

	var defIntVal int
	var defStrVal string
	var defBoolVal bool
	var defFloat64Val float64
	var defFloat32Val float32

	var (
		name      = "Nihanth"
		age       = 25
		role      = "Software Engineer"
		interests = "Golang"
	)

	fmt.Println("Default values", defIntVal, defStrVal, defBoolVal, defFloat64Val, defFloat32Val)
	fmt.Println("User Details", name, age, role, interests)
}
