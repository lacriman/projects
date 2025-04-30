package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	tasks := [][2]int{
		{10, 2},
		{9, 3},
		{5, 0},
		{100, 10},
	}

	for _, pair := range tasks {
		go divideAndPrint(pair[0], pair[1])
	}
	time.Sleep(1 * time.Second)
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
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
