package main

import (
	"fmt"
	"sync"
	"time"
)

func main_waitGroup() {
	fmt.Println("Starting main Goroutine")
	var waitgroup sync.WaitGroup
	waitgroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(waitgroup *sync.WaitGroup, number int) {
			defer waitgroup.Done()
			fmt.Printf("Starting routine %d \n", number)
			time.Sleep(time.Second)
			fmt.Printf("Done with routine %d \n", number)
		}(&waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}

// Output :
// Starting main Goroutine
// Starting routine 4
// Starting routine 0
// Starting routine 1
// Starting routine 2
// Starting routine 3
// Done with routine 0
// Done with routine 4
// Done with routine 1
// Done with routine 3
// Done with routine 2
// Finishing main Goroutine

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func routine(waitgroup *sync.WaitGroup, number int) {
	defer waitgroup.Done()
	fmt.Printf("Starting routine %d \n", number)
	time.Sleep(time.Second)
	fmt.Printf("Done with routine %d \n", number)
}

func main() {
	fmt.Println("Starting main Goroutine")
	var waitgroup sync.WaitGroup

	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go routine(&waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}
*/
// Output :
// Starting main Goroutine
// Starting routine 4
// Starting routine 2
// Starting routine 3
// Starting routine 0
// Starting routine 1
// Done with routine 1
// Done with routine 0
// Done with routine 3
// Done with routine 4
// Done with routine 2
// Finishing main Goroutine
