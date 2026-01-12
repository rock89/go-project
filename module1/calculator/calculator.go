package main

import (
	"fmt"
)

// Function for Addition
func add(a, b float64) float64 {
	return a + b
}

// Function for Subtraction
func subtract(a, b float64) float64 {
	return a - b
}

// Function for Multiplication
func multiply(a, b float64) float64 {
	return a * b
}

// Function for Division
func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero is not allowed")
		return 0
	}
	return a / b
}

func main() {
	var choice int
	var num1, num2 float64

	for {
		fmt.Println("\n===== Calculator Menu =====")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Exiting calculator...")
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		switch choice {
		case 1:
			fmt.Println("Result:", add(num1, num2))
		case 2:
			fmt.Println("Result:", subtract(num1, num2))
		case 3:
			fmt.Println("Result:", multiply(num1, num2))
		case 4:
			fmt.Println("Result:", divide(num1, num2))
		}
	}
}
