package main

import "fmt"

func main() {
	var userInput int
	num1 := 0
	num2 := 1

	fmt.Println("It's a fibonacci sequence, how many numbers do you want?: ")
	fmt.Scan(&userInput)

	for i := 0; i < userInput; i++ {
		fmt.Print(num1, " ")
		next := num1 + num2
		num1 = num2
		num2 = next
	}
}
