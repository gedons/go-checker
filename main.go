package main

import (
	"fmt"
)

func main() {
	 input := 1

	fmt.Printf("======================== Welcome To DSPG Microsoft Imagine Academy Student Result Checker ========================\n")
	fmt.Print("Select (1) to login as a Student\n")
	fmt.Print("Select (2) to login as a Teacher\n")

	switch input {

	case 1:
		studentLogin()
	case 2:
		teacherLogin()
	default:
		fmt.Println("Invalid Selection")

	}
}

//LOGIN FUNCTION
func studentLogin() {
	fmt.Println("student")
}

func teacherLogin() {
	fmt.Println("teachers")
}
