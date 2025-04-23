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

			if n != 1 || err != nil || usersGrade <= 0 {
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
	maxGrade := grades[0]
	minGrade := grades[0]
	gradesSum := 0
	
	// Find an average grade
	for i := 0; i < len(grades); i++ {
		gradesSum += grades[i] 
		averageGrade = gradesSum / len(grades)
	}
	
	// Find the max grade
	for _, grade := range grades[1:] {
		if grade > maxGrade {
			maxGrade = grade
		}
	}
	
	// Find the min grade
	for _, grade := range grades[1:] {
		if grade < minGrade {
			minGrade = grade
		}
	}

	fmt.Printf("Your average grade is: %d, your max grade is: %d, and your lowest grade is: %d", averageGrade, maxGrade, minGrade)
}
