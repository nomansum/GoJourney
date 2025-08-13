package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	fmt.Println("Sending a string into channel ")
	ch <- "hello"
	fmt.Println("String has been retrieved from channel>... ")
}

// ---getting data from the channel---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String received from sendData is ", <-ch)

}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
}
