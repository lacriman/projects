package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	apiList := []string{"api1", "api2", "api3", "api4", "api5", "api6"}
	var wg sync.WaitGroup
	start := time.Now()

	for _, api := range apiList {
		wg.Add(1)
		go func(api string) {
			defer wg.Done()
			pingApi(api)
		}(api)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("All API pinged in %v\n", elapsed)
}

func pingApi(api string) {
	fmt.Printf("Pinging %s ....\n", api)
	delay := time.Duration(rand.Intn(400)+100) * time.Millisecond
	time.Sleep(delay)
	fmt.Printf("%s is responded in %v\n", api, delay)
}
