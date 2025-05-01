package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <word to find> file1.txt file2.txt ...")							
		return
	}

	wordToFind:= os.Args[1]
	files:= os.Args[2:]
	totalCount:= 0

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, file := range files {
		file:= file
		wg.Add(1)
		go func() {
			defer wg.Done()
			count, err:= countWordInFile(file, wordToFind)
			if err != nil {
			  fmt.Printf("Error reading %s: %v\n", file, err)
			  return
			}

			mu.Lock()
			totalCount += count
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("The word %s appeared %d times across all files.\n", wordToFind, totalCount)
}

func countWordInFile(filePath string, word string) (int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		if strings.EqualFold(scanner.Text(), word) {
			count++
		}
	}
	return count, scanner.Err()
}