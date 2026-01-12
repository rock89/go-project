package main

import "fmt"

func main() {
	var students []Student
	var choice int

	for {
		showMenu()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addStudent(&students)
		case 2:
			viewStudents(students)
		case 3:
			searchStudent(students)
		case 4:
			updateStudent(&students)
		case 5:
			deleteStudent(&students)
		case 6:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
