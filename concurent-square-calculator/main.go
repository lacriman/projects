package main

import (
	"fmt"
	"sync"
)

func main() {
	results := make(map[int] int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			square := n * n

			mu.Lock()
			results[n] = square
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Results: ", results)
}

