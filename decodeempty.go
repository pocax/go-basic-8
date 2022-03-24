package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
		{
			"fullname": "John Doe",
			"email": "john@doe.com",
			"age": 30
		}
	`

	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &temp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = temp.(map[string]interface{})
	fmt.Println("fullname:", result["fullname"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
