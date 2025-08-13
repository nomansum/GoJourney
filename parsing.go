package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var mp = make(map[string]interface{})

const l = 12

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type myType float64

func (mytype myType) sumWithNum(a float64) (result float64) {
	result = float64(mytype) + a
	return
}

func main_parsing() {
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

	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Jo", Age: 24}
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Can not create file")
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("FROM HERE")
	var a = myType(23.023)
	fmt.Println(a)
	fmt.Println(a.sumWithNum(1224))

}
