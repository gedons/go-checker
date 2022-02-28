package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/TwiN/go-color"
)

func main() {

	fmt.Printf("\n\t\t================== Welcome To DSPG Microsoft Imagine Academy Student Result Checker ================== \n\n")
	fmt.Printf("\t Input (1) to login as a STUDENT\n\n")
	fmt.Printf("\t Input (2) to login as a TEACHER\n\n")

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
		time.Sleep(2 * time.Second)
		main()
	}
}

//LOGIN FUNCTIONs
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
			println(color.Ize(color.Green, "LOGGING IN........."))
			time.Sleep(3 * time.Second)
			CallClear()
			teacherWindow()
			break

		} else {
			println(color.Ize(color.Red, "Invalid Details"))
			continue
		}
	}
}

//teacher window function
func teacherWindow() {
	fmt.Printf("\n\t\t Follow the instructions below!!!!\n\n")
	fmt.Printf("\t 1 To Insert Student Details\n\n")
	fmt.Printf("\t 2 To View Student Details\n\n")
	fmt.Printf("\t 3 To Delete Student Details\n\n")
	fmt.Printf("\t 4 To Logout\n\n")

	var teacherInput int
	fmt.Scan(&teacherInput)

	switch teacherInput {
	case 1:
		teacherWindowInsert()
	case 2:
		teacherWindowView()
	case 3:
		teacherWindowDelete()
	case 4:
		teacherLogin()
	default:
		fmt.Println("Invalid Selection")
		time.Sleep(2 * time.Second)
		teacherWindow()
	}

}

//teachers window to insert new student
func teacherWindowInsert() {
	var fname string
	var lname string
	var score1 int
	var score2 int
	var total int

	fmt.Println("Input Student Firstname : ")
	fmt.Scan(&fname)
	var path = fname + ".txt"

	fmt.Println("Input Student Lastname : ")
	fmt.Scan(&lname)

	fmt.Println("Input First Semester Score : ")
	fmt.Scan(&score1)

	fmt.Println("Input Second Semester Score : ")
	fmt.Scan(&score2)

	total = score1 + score2

	// check if file exists
	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	//WRITING STUDENT DETAILS TO A FILE
	//first open the file
	var file1, err1 = os.OpenFile(path, os.O_RDWR, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file1.Close()

	//then write some texts line by line
	_, err1 = file1.WriteString("First Name : " + fname + "\n")
	if err1 != nil {
		log.Fatal(err1)
	}

	_, err1 = file1.WriteString("Last Name : " + lname + "\n")
	if err1 != nil {
		log.Fatal(err1)
	}

	_, err1 = file1.WriteString(fmt.Sprintf("Score 1 : %d\n", score1))
	if err1 != nil {
		log.Fatal(err1)
	}

	_, err1 = file1.WriteString(fmt.Sprintf("Score 2 : %d\n", score2))
	if err1 != nil {
		log.Fatal(err1)
	}

	_, err1 = file1.WriteString(fmt.Sprintf("Total Score : %d\n", total))
	if err1 != nil {
		log.Fatal(err1)
	}

	// Save file changes.
	err1 = file1.Sync()
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("Student Details Created Successfully (%v)\n", path)

	println(color.Ize(color.Green, "REDIRECTING TO HOME......."))
	time.Sleep(3 * time.Second)
	CallClear()
	teacherWindow()

}


func teacherWindowView() {

	var path string

	fmt.Println("\t Available Student Records")
	//lists all the files in the directory
	files, err := os.ReadDir("./")
	if err != nil{
		fmt.Println(err.Error())
	}
	if files == nil {
		fmt.Println("No Records Available!!!")
		time.Sleep(3 * time.Second)
		teacherWindow()
	}

	// returns all the files with a .txt extension
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".txt" {
			fmt.Println(f.Name())
		}

	}

	fmt.Println("\t Input A Student Name Above To View Details")
	fmt.Scan(&path)

	// Open file for reading.
	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if err2 != nil {
		println(color.Ize(color.Red, "INVALID FILE SELECTED"))
		time.Sleep(2 * time.Second)
		teacherWindow()
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err2 = file.Read(text)

		// Break if finally arrived at end of file
		if err2 == io.EOF {
			break
		}
		// Break if error occured
		if err2 != nil && err2 != io.EOF {
			fmt.Println(err2.Error())
			break
		}
	}

	//display file
	println(color.Ize(color.Green, "OPENING RECORD, PLEASE WAIT...."))
	time.Sleep(3 * time.Second)
	fmt.Println(string(text))
	teacherWindow()

}


func teacherWindowDelete() {
	
	var path string

	fmt.Println("\t Available Student Records")
	//lists all the files in the directory
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(2 * time.Second)
		teacherWindow()
	}

	// returns all the files with a .txt extension
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".txt" {
			fmt.Println(f.Name())
		}

	}

	fmt.Println("\t Input A Student Name Above To Delete Details")
	fmt.Scan(&path)

	 // delete file
	 var errs = os.Remove(path)
	 if errs != nil {
		fmt.Println(errs.Error())
		time.Sleep(2 * time.Second)
		teacherWindow()
	}

	println(color.Ize(color.Red, "DELETING....."))
	time.Sleep(3 * time.Second)
	fmt.Println(color.Ize(color.Green, "DELETED SUCCESSFULLY!!!"))
	time.Sleep(2 * time.Second)
	teacherWindow()
}

// func studentWindow() {
// 	fmt.Println("student window")
// }

//CODES TO CLEAR THE TERMINAL SCREEN
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
