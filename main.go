package main

import (
	"fmt"
	"github.com/TwiN/go-color"
)


func main() {

	fmt.Printf("\n\t\t================== Welcome To DSPG Microsoft Imagine Academy Student Result Checker ================== \n")
	fmt.Printf("Input (1) to login as a STUDENT\n")
	fmt.Printf("Input (2) to login as a TEACHER\n")

	//get the user input in the switch case for the two available functions so far
	var input int
	fmt.Scan(&input)

	switch input {

	case 1:
		studentLogin()

	case 2:
		teacherLogin()

	default:
		fmt.Printf("Invalid Selection \n")
	}
}

//LOGIN FUNCTION
func studentLogin() {
	fmt.Println("student")
}

func teacherLogin() {

	username := "admin"
	password := "12345"
	var iuname string
	var ipass string

	fmt.Printf("\n\t\t================== Teacher Login ================== \n")

	for {
		fmt.Println("Input Username : ")
		fmt.Scan(&iuname)

		fmt.Println("Input Password : ")
		fmt.Scan(&ipass)

		if iuname == username && ipass == password {
			teacherWindow()
			break

		} else {
			fmt.Println("Invalid Details")
			continue
		}
	}
}

func teacherWindow() {
	fmt.Println("teacher window")
}
