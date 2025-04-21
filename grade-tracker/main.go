package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the grade tracker! Write your 5 grades below")
	grades := []int{}

	for i := 1; i <= 5; i++ {
		usersGrade := 0
		for {
			fmt.Printf("Grade number %d: ", i)
			n, err := fmt.Scan(&usersGrade)

			if n != 1 || err != nil || usersGrade < 0 {
				fmt.Println("Invalid input! Please enter a positive integer only.")

				var junk string 
				fmt.Scanln(&junk)
				continue
			}

			grades = append(grades, usersGrade)
			break
		}
	}
	averageGrade := 0
	// maxGrade := 0
	// minGrade := 0
	gradesSum := 0
	for i := 0; i < len(grades); i++ {
		gradesSum += grades[i] 
		averageGrade = gradesSum / len(grades)
	}
	fmt.Println(averageGrade)
}
