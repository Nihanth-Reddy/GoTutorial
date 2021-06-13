package main

import "fmt"

type User struct {
	Name    string
	Email   string
	age     int
	courses []string
}

func main() {

	tic := User{"Tic", "tic@game.com", 19, []string{"Golang", "Javascript"}}

	var tac = new(User)
	tac.Name = "Tac"
	tac.age = 20
	tac.Email = "tac@game.com"
	tac.courses = []string{"Python", "Java"}

	var toe = &User{"Toe", "toe@game.com", 22, []string{"C", "C++"}}

	if contains(tic.courses, "C") {
		fmt.Println("Tic has enrolled for C course")
	} else {
		fmt.Println("Tic hasn't enrolled for C course, adding him forcefully")
		tic.courses = append(tic.courses, "C")
	}
	if contains(tac.courses, "C") {
		fmt.Println("Tac has enrolled for C course")
	} else {
		fmt.Println("Tac hasn't enrolled for C course, adding him forcefully")
		tac.courses = append(tac.courses, "C")
	}
	if contains(toe.courses, "C") {
		fmt.Println("Toe has enrolled for C course")
	} else {
		fmt.Println("Toe hasn't enrolled for C course, adding him forcefully")
		toe.courses = append(toe.courses, "C")
	}

}

func contains(ss []string, s string) bool {
	for _, val := range ss {
		if val == s {
			return true
		}
	}
	return false
}
