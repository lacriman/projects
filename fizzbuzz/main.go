package main

import (
	"fmt"
)

// func main() {
// 	numbers := 0
// 	fmt.Println("Welcome to the fizzbuzz!\nHow many numbers should we run FizzBuzz on?: ")
// 	fmt.Scan(&numbers)

// 	for numbers <= 0 {
// 		fmt.Println("You have you write an integer that greater then 0:")
// 		fmt.Scan(&numbers)
// 	}

// 	for i := 1; i <= numbers; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Print("FizzBuzz ")
// 		} else if i%3 == 0 {
// 			fmt.Print("Fizz ")
// 		} else if i%5 == 0 {
// 			fmt.Print("Buzz ")
// 		} else {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }

func main() {
	for {
		fmt.Println("Welcome to the fizzbuzz!")

		numbers := getValidInt("How many numbers should we run FizzBuzz on?: ")
		fizz := getValidInt("Enter your Fizz number: ")
		buzz := getValidInt("Enter your Buzz number: ")

		for i := 1; i <= numbers; i++ {
			if i%fizz == 0 && i%buzz == 0 {
				fmt.Print("FizzBuzz ")
			} else if i%fizz == 0 {
				fmt.Print("Fizz ")
			} else if i%buzz == 0 {
				fmt.Print("Buzz ")
			} else {
				fmt.Print(i, " ")
			}
		}
		
		fmt.Println("\nWould you like to play again? (y/n):")
		playAgain := "n"
		fmt.Scan(&playAgain)
		if playAgain == "n" {
			break
		}
	}
}

func getValidInt(message string) int {
	fmt.Println(message)
	num := 0
	fmt.Scan(&num)
	if num <= 0 {
		fmt.Println("You have you write an integer that greater then 0:")
		fmt.Scan(&num)
	}

	return num
}
