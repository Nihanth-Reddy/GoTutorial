package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Student struct {
	Name   string
	Grades []int
	Age    int
}

// Here we are not gonna update any Struct value hence we are passing struct directly
func (s Student) getAverageGrades() float64 {
	sum := 0
	for _, val := range s.Grades {
		sum += val
	}
	return float64(sum / len(s.Grades))
}

// Here we are gonna update age by 1, for changes to get reflected we are passing pointer to struct
func (s *Student) incrementAge() {
	s.Age += 1
}

func (s Student) getMaxGrade(prefix string) string {
	curr_max := s.Grades[0]
	for i := 1; i < len(s.Grades); i++ {
		if s.Grades[i] > curr_max {
			curr_max = s.Grades[i]
		}
	}
	return prefix + strconv.Itoa(curr_max)
}
func main() {
	fmt.Println("Welcome to Struct methods")
	s1 := Student{"Tic", []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}, 19}
	s2 := Student{"Tac", []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}, 20}
	s3 := Student{"Toe", []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}, 21}

	fmt.Println("Grades of Student s1 are: ", s1.Grades)
	fmt.Println("Average grades of Student s1 is: ", s1.getAverageGrades())

	fmt.Println("Age of Student s2 is: ", s2.Age)
	s2.incrementAge()
	fmt.Println("Incremented age of Student s2 is: ", s2.Age)

	fmt.Println("Grades of Student s3 are: ", s3.Grades)
	fmt.Println("Max grade of Student s3 is: ", s3.getMaxGrade("Max grade of "+s3.Name+"is :"))

}
