package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var balance int
	fmt.Println("Welcome to the ATM")
	checkPin()

	for {
		input := promptUser("\nWrite your balance: ")
		bal, err := readPositiveInt(input)

		if err != nil {
			fmt.Printf("\nError: %v", err)
			continue
		}
		balance = bal
		fmt.Printf("\nYour balance is: %d", balance)
		break
	}

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("\t1 = Withdraw\n\t2 = Deposit\n\t3 = Exit")
		input := promptUser("\nAction: ")
		action, err := readPositiveInt(input)

		if err != nil {
			fmt.Printf("\nError: %v", err)
			continue
		}

		switch action {
		case 1:
			fmt.Println("\nWithdraw selected.")
			err := withdraw(&balance)
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
			}
		case 2:
			fmt.Println("\nDeposit selected.")
			
			showBalance(balance)
			err := deposit(&balance)
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
			}
		case 3:
			fmt.Println("\nGoodbye.")
			os.Exit(0)
		default:
			fmt.Println("You have to write a number from 1 to 3.")
			continue

		}
	}
}

func checkPin() {
	pin := 1234
	attempts := 3

	for attempts > 0 {
		input := promptUser("\nPlease write your PIN code: ")
		guess, err := readPositiveInt(input)

		if err != nil {
			fmt.Printf("\nError: %v", err)
			continue
		}

		if guess == pin {
			fmt.Println("PIN correct.")
			break
		}
		attempts--

		if attempts == 0 {
			fmt.Println("\nYou have no attempts left, the card is blocked.")
			os.Exit(1)
		} else {
			fmt.Printf("\nWrong! You have %d more attempts. \n", attempts)
		}

	}

}

func promptUser(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)

	

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readPositiveInt(input string) (int, error) {
	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, fmt.Errorf("Something went wrong: %v", err)
	}
	if num <= 0 {
		return 0, errors.New("The number has to be greater than 0")
	}
	return num, nil
}

func withdraw(balance *int) error {
	if *balance == 0 {
		return errors.New("Your balance is 0, you can't withdraw.")
	}

	for {
		input := promptUser("\nHow much money do you want to withdraw?\nAmount: ")
		amount, err := readPositiveInt(input)
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			continue
		}

		if amount > *balance {
			fmt.Println("Sorry, you don't have enough money.")
			continue
		}

		*balance -= amount
		showBalance(*balance)
		return nil
	}
}

func deposit(balance *int) error {
	input := promptUser("\nHow much money do you want to deposit?\nAmount: ")
	amount, err := readPositiveInt(input)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	*balance += amount
	return nil
}

func showBalance(balance int) {
	fmt.Printf("\nYour balance now is: %d", balance)
}
