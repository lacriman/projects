package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var balance int
	var action int
	fmt.Println("Welcome to the ATM")

	checkPin()
	balance = readPositiveInt("Write your balance: ")
	fmt.Printf("Your balance is: %d\n", balance)

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("\t1 = Withdraw\n\t2 = Deposit\n\t3 = Exit")
		action = readPositiveInt("Action: ")

		switch action {
		case 1:
			fmt.Println("\nWithdraw selected.")
			withdraw(&balance)
			break
		case 2:
			fmt.Println("\nDeposit selected.")
			deposit(&balance)
			showBalance(balance)
			break

		case 3:
			fmt.Println("\nGoodbye.")
			os.Exit(0)
		default:
			fmt.Println("\nWrong input, write something from 1 to 3.")

		}
	}

}

func checkPin() {
	pin := 1234
	attempts := 3

	for attempts > 0 {
		guess := readPositiveInt("Please write your PIN code: ")

		if guess == pin {
			fmt.Println("PIN correct.")
			return
		}
		attempts--
		if attempts == 0 {
			fmt.Println("You have no attempts left, the card is blocked.")
			os.Exit(1)
		} else {
			fmt.Printf("Wrong! You have %d more attempts. \n", attempts)
		}

	}

}

func readPositiveInt(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)

		if err != nil || num <= 0 {
			fmt.Println("Wrong input, please write a positive number.")
			continue
		}

		return num
	}
}

func withdraw(balance *int) {
	if *balance == 0 {
		fmt.Println("\nYour balance is 0, you can't withdraw.")
		return
	}
	for {
		amount := readPositiveInt("\nHow much money do you want to withdraw?\nAmount: ")
		if *balance-amount < 0 {
			fmt.Println("\nSorry, you don't have enough money.")
			continue
		}
		*balance -= amount
		showBalance(*balance)
		break
	}
}

func deposit(balance *int) {
	amount := readPositiveInt("\nHow much money do you want to deposit?\nAmount: ")
	*balance += amount
}

func showBalance(balance int) {
	fmt.Printf("\nYour balance now is: %d", balance)
}
