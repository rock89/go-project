package main

import "fmt"

// Add Student
func addStudent(students *[]Student) {
	var s Student

	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&s.ID)

	fmt.Print("Enter Name: ")
	fmt.Scanln(&s.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&s.Age)

	fmt.Print("Enter Course: ")
	fmt.Scanln(&s.Course)

	fmt.Print("Enter Marks: ")
	fmt.Scanln(&s.Marks)

	*students = append(*students, s)
	fmt.Println("Student added successfully.")
}

// View All Students
func viewStudents(students []Student) {
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, s := range students {
		fmt.Println("----------------------")
		fmt.Println("ID:", s.ID)
		fmt.Println("Name:", s.Name)
		fmt.Println("Age:", s.Age)
		fmt.Println("Course:", s.Course)
		fmt.Println("Marks:", s.Marks)
	}
}

// Search Student by ID
func searchStudent(students []Student) {
	var id int
	fmt.Print("Enter Student ID to search: ")
	fmt.Scanln(&id)

	for _, s := range students {
		if s.ID == id {
			fmt.Println("Student Found:")
			fmt.Println("Name:", s.Name)
			fmt.Println("Age:", s.Age)
			fmt.Println("Course:", s.Course)
			fmt.Println("Marks:", s.Marks)
			return
		}
	}
	fmt.Println("Student not found.")
}

// Update Student
func updateStudent(students *[]Student) {
	var id int
	fmt.Print("Enter Student ID to update: ")
	fmt.Scanln(&id)

	for i, s := range *students {
		if s.ID == id {
			fmt.Print("Enter New Name: ")
			fmt.Scanln(&(*students)[i].Name)

			fmt.Print("Enter New Age: ")
			fmt.Scanln(&(*students)[i].Age)

			fmt.Print("Enter New Course: ")
			fmt.Scanln(&(*students)[i].Course)

			fmt.Print("Enter New Marks: ")
			fmt.Scanln(&(*students)[i].Marks)

			fmt.Println("Student updated successfully.")
			return
		}
	}
	fmt.Println("Student not found.")
}

// Delete Student
func deleteStudent(students *[]Student) {
	var id int
	fmt.Print("Enter Student ID to delete: ")
	fmt.Scanln(&id)

	for i, s := range *students {
		if s.ID == id {
			*students = append((*students)[:i], (*students)[i+1:]...)
			fmt.Println("Student deleted successfully.")
			return
		}
	}
	fmt.Println("Student not found.")
}
