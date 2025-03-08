package syncWaitGroup

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that the goroutine is done
	fmt.Println("\nWorker", i, "started")
	//
	fmt.Println("Worker", i, "ending")
}

func SyncWaitGroup() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment the wait group counter
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("\nWorker Loop Completed")
}

// {

// 		// channel example
// 		{
// 			ch := make(chan int)
// 			for j := 1; j < 3; j++ {
// 				// set channel
// 				ch <- j
// 				go workerCh(j, ch)
// 			}

// 			fmt.Println("\nWorker using Channel Loop Completed")
// 		}
// 	func workerCh(i int, ch chan int) {
// 		fmt.Println("Worker", i, "started")
// 		ch <- i
// 		fmt.Println("Worker", i, "ending")
// 	}
// }
