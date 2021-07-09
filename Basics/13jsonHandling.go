package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type SensorReading struct {
	Name     string     `json:"name"`
	Capacity int        `json:"capacity"`
	Time     string     `json:"time"`
	Info     SensorInfo `json:"info"`
}

type SensorInfo struct {
	Description string `json:"desc"`
}

func main() {
	// Encoding / Marshal :- Encodes struct to json string
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	// Decoding / Unmarshal :- Decodes json string to struct
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z", "info": { "desc": "Defines battery stats"}}`

	var reading SensorReading
	err = json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)
}
