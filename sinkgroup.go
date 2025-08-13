package main

import (
	"fmt"
	"sync"
	"time"
)

var arrTest = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var minicache sync.Map

func towrite(text string) {

	for _, v := range arrTest {
		minicache.Store(v, text)
		time.Sleep(10 * time.Millisecond)
	}

}

func main_sinkGroup() {
	go towrite("test1")
	go towrite("test2")
	go func() {
		time.Sleep(10 * time.Millisecond)
		for _, v := range arrTest {
			if val, ok := minicache.Load(v); ok {
				fmt.Println(v, val)

			} else {
				fmt.Println(v, "not get")
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	time.Sleep(10 * time.Second)
}

/*
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	counter++
}
*/

/*

var i = 10
mutex := &sync.RWMutex{}
mutex.Lock()
// only one Goroutine can access this code at a time
i++
mutex.Unlock()
mutex.RLock()
i++ // concurrent reads
mutex.RUnlock()

*/
