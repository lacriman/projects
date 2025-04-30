package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	tasks := [][2]int{
		{10, 2},
		{9, 3},
		{5, 0},
		{100, 10},
	}

	for _, pair := range tasks {
		pair := pair
		wg.Add(1)
		go func() {
			defer wg.Done()
			divideAndPrint(pair[0], pair[1])
		}()
	}
	wg.Wait()
}

func divideAndPrint(a, b int) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot divide %v and %v. Error: %v\n", a, b, err)
		return
	}

	fmt.Printf("%v / %v = %v\n", a, b, result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}
