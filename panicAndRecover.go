package main

import "fmt"

func main() {
	fmt.Println("main function start")
	defer fmt.Println("main function er  defer")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic dhora porse:", r)
		}
	}()
	riskyFunction()
	fmt.Println("main function er end") // এটা এক্সিকিউট হবে যদি প্যানিক রিকভার হয়
}

func riskyFunction() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic dhora porse ", r)
	// 	}
	// }()
	defer fmt.Println("riskyFunction er defer")
	fmt.Println("riskyFunction cholche ")
	panic("kichu vul hoyeche!")
	fmt.Println("eta kokhono print hbe na ")
}
