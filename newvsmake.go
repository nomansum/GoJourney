package main

import "fmt"

func main_newVsMake() {
	p := new(int)   // *int type
	fmt.Println(*p) // 0 (zero value of int)
	*p = 42
	fmt.Println(*p) // 42

	s := new([]int) // *[]int type
	fmt.Println(*s) // nil (slice zero value)
	// Slice

	*s = make([]int, 0, 0)
	*s = append((*s), 12)
	fmt.Println((*s)[0])

}
