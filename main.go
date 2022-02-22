package main

import (
	"fmt"
)

func main() {
	 var input int
	 fmt.Scanln(input)

 

		switch input {

		case 1:
			fmt.Println("Winter.")

		case 2:
			fmt.Println("Autumn.")
	}
}

//LOGIN FUNCTION
func studentLogin() {
	fmt.Println("student")
}

func teacherLogin() {
	fmt.Println("teachers")
}
