package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	files := os.Args[1:]
	
	for _, file := range files {
		file:= file
		wg.Add(1)
		go func() {
			defer wg.Done()
			err:=lineCounter(file)
			if err != nil {
			  fmt.Printf("error: %v\n", err)
			}
		}()
	}

	wg.Wait()
}

func lineCounter(file string) error{
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Can't find a file %s", file)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Number of lines in file %s: %d\n",file, len(lines))
	return	nil
}
