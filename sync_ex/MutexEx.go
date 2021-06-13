package sync_ex

import (
	"fmt"
	"runtime"
	"sync"
)

func MutexEx() {
	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		defer mu.Unlock()
		mu.Lock()
		balance += amount
	}

	withdrawal := func(amount int) {
		defer mu.Unlock()
		mu.Lock()
		balance -= amount
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check result.

	// fix the issue for consistent output.
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
