package main

import "fmt"

func showMenu() {
	fmt.Println("\n===== Student Management System =====")
	fmt.Println("1. Add Student")
	fmt.Println("2. View All Students")
	fmt.Println("3. Search Student by ID")
	fmt.Println("4. Update Student")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Exit")
	fmt.Print("Enter your choice: ")
}
