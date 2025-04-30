package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	files := os.Args[1:]

	for _, file := range files {
		file := file
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := lineCounter(file)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}()
	}

	wg.Wait()
}

func lineCounter(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("cannot read file %v: %v\n",file, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		counter++
	}
	fmt.Printf("The number of lines in %s is %d\n", file, counter)
	return nil
}
