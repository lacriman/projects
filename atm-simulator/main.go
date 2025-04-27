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

	err := checkPin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	enterBalance(&balance)
	action(&balance)
}

func checkPin() error {
	pin := 1234
	attempts := 3

	for attempts > 0 {
		guess, err := readPositiveInt("\nPlease write your PIN code: ")

		if err != nil {
			attempts--
			fmt.Fprintf(os.Stderr, "Invalid input! Please write numbers only. You have %d more attempts.\n", attempts)
		} else if guess == pin {
			fmt.Println("PIN correct.")
			return nil
		} else {
			attempts--
			fmt.Fprintf(os.Stderr, "\nWrong! You have %d more attempts. \n", attempts)
		}
	}
	if attempts == 0 {
		fmt.Fprintf(os.Stderr, "\nYou have no attempts left, the card is blocked.\n")
		return errors.New("no attempts left, card blocked")
	}

	return nil
}

func promptUser(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readPositiveInt(message string) (int, error) {
	input := promptUser(message)
	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}
	if num <= 0 {
		return 0, errors.New("The number has to be greater than 0")
	}
	return num, nil
}

func enterBalance(balance *int) error {
	bal, err := readPositiveInt("\nWrite your balance: ")
	if err != nil {
		return err
	}
	*balance = bal
	showBalance(*balance)
	return nil
}

func action(balance *int) error {
	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("\t1 = Withdraw\n\t2 = Deposit\n\t3 = Exit")
		action, err := readPositiveInt("\nAction: ")

		if err != nil {
			return err
		}

		switch action {
		case 1:
			fmt.Println("\nWithdraw selected.")
			err := withdraw(balance)
			if err != nil {
				fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			}
		case 2:
			fmt.Println("\nDeposit selected.")
			err := deposit(balance)
			if err != nil {
				fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
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

func withdraw(balance *int) error {
	amount, err := readPositiveInt("\nHow much money do you want to withdraw?\nAmount: ")

	if err != nil {
		return err
	}
	if *balance == 0 {
		return errors.New("Your balance is 0, you can't withdraw.")
	}
	if amount > *balance {
		return errors.New("Sorry, you don't have enough money.")
	}

	*balance -= amount
	showBalance(*balance)

	return nil
}

func deposit(balance *int) error {
	amount, err := readPositiveInt("\nHow much money do you want to deposit?\nAmount: ")
	if err != nil {
		return err
	}
	*balance += amount
	showBalance(*balance)

	return nil
}

func showBalance(balance int) {
	fmt.Printf("\nYour balance now is: %d\n", balance)
}
