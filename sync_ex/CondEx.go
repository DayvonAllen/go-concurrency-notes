package sync_ex

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func CondEx() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(2)
	go func() {
		defer wg.Done()
		c.L.Lock()
		// suspend goroutine until sharedRsc is populated.
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	// writes changes to sharedRsc
	go func() {
		defer wg.Done()
		c.L.Lock()
		sharedRsc["rsc1"] = "foo"
		c.Signal()
		c.L.Unlock()
	}()

	wg.Wait()
}

