package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
		{
			"fullname": "John Doe",
			"email": "john@doe",
			"age": 30
		}
	`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("fullname:", result["fullname"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
