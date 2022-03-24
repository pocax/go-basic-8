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
		[{
			"fullname": "John Doe",
			"email": "john@doe.com",
			"age": 30
		},
		{
			"fullname": "Jane Doe",
			"email": "linus@doe.com",
			"age": 25
		}]
	`
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

}
