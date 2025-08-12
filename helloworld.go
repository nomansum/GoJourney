package main

import (
	"fmt"
	"time"
)

func worker(cancel <-chan struct{}) {
	fmt.Println("Worker started")
	defer fmt.Println("Worker stopped")
	for {
		select {
		case <-cancel:
			fmt.Println("Worker canceled")
			return
		default:
			// Do some work here
			time.Sleep(time.Second)
			fmt.Println("Working...")
		}
	}
}

func main() {
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5 * time.Second)
	fmt.Println("Canceling worker")
	close(cancel)
	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine stopped")
}

/*
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func worker(cancel <-chan struct{}) {
	fmt.Println("Worker started")
	defer fmt.Println("Worker Stopped")
	for {
		select {
		case <-cancel:
			fmt.Println("Worker cancelled")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Working")
		}
	}
}
func main() {

	cancel := make(chan struct{})

	go worker(cancel)

	go numbers()
	go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

*/
