package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	totalErrorLines := 0

	fileNames := os.Args[1:]
	for _, file := range fileNames {
		file := file
		wg.Add(1)
		go func() {
			defer wg.Done()
			errorLines, err := countErrors(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot read %s: %v\n", file, err)
			}
			mu.Lock()
			defer mu.Unlock()
			totalErrorLines += errorLines
		}()
	}
	wg.Wait()
	fmt.Println(totalErrorLines)
}

func countErrors(filePath string) (int, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorLines := 0

	for scanner.Scan() {
		errorLines++
	}
	return errorLines, nil
}
