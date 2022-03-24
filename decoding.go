package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
		{
			"fullname": "John Doe",
			"email": "john@doe.com",
			"age": 30
		}
	`

	var result Employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("fullname:", result.Fullname)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}
