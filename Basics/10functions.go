package main

import "fmt"

func main() {
	defer printFirst()
	fmt.Println("This is the main function")
	oddNumber := oddNumberGenerator()
	fmt.Println(oddNumber())
	fmt.Println(oddNumber())
	fmt.Println("Average of 10, 12, 15 is: ", average(10, 12, 15))
	fmt.Println("Average of 3, 6, 9 is: ", average(3, 6, 9))
	fmt.Println("is 4 a Prime Number: ", isPrime(4))
	fmt.Println("is 2 a Prime Number: ", isPrime(2))
	printSecond()

	fmt.Println(contains([]string{"Nihanth", "Reddy", "Kottam"}, "Nihanth"))
	fmt.Println(contains([]string{"Nihanth", "Reddy", "Kottam"}, "Ajay"))

}

// Closure functions -> function returning function
func oddNumberGenerator() func() uint {
	i := uint(1)
	return func() (num uint) {
		num = i
		i += 2
		return
	}
}

// Variadic functions -> takes special type of argument in the last argument
func average(nums ...int) float64 {
	sum := 0
	for _, val := range nums {
		sum = val + sum
	}
	return float64(sum / len(nums))
}

// Verify if given number is prime or not
func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("Given number is not prime as it is divisible by ", i)
			return false
		}
	}
	return true
}

func printFirst() {
	fmt.Println("This the the first function")
}

func printSecond() {
	fmt.Println("This the the second function")
}

func contains(ss []string, s string) string {
	for _, x := range ss {
		if x == s {
			return "Exists in the slice"
		}
	}
	panic("String doesn't exist in the slice")
}
