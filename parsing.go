package main

import (
	"encoding/json"
	"fmt"
)

var mp = make(map[string]interface{})

const l = 12

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	type response2 struct {
		Page   int      `json:"page-number"`
		Fruits []string `json:"fruits-list"`
	}
	str := `
	{
	 "number":1,
	 "fruits-list":["apple","peach"]
	}

	`
	res := &response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	fmt.Println("AFter This ignore")

	p := Person{Name: "Alice", Age: 25}

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes:", jsonBytes)          // Output: [123 34 110 97 109 101 34 58 34 65 108 105 99 101 34 44 34 97 103 101 34 58 50 53 125]
	fmt.Println("String:", string(jsonBytes)) // Output: {"name":"Alice","age":25}
}
