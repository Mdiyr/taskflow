package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type Task struct {
	ID         int
	UserID     int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
}
type Category struct {
	ID     int
	UserID int
	Title  string
	Color  string
}

var userStorage []User
var AuthenticatedUser *User
var taskStorage []Task
var categoryStorage []Category

const userStoragePath = "user.txt"

func main() {

	fmt.Println("Hello to TODO app")

	//load user storage from file
	loadUserStorageFromFile()

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()

	}
}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && AuthenticatedUser == nil {
		fmt.Println("You must log in first!")
		login()

		if AuthenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "list-task":
		listTask()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}

}
func createTask() {

	fmt.Println("Welcome: ", AuthenticatedUser.Email)
	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("categoryID is not valid integer: %v\n", err)
		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == AuthenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Printf("categoryID is not found!\n")
		return
	}

	task := Task{
		ID:         len(taskStorage) + 1,
		UserID:     AuthenticatedUser.ID,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
	}
	taskStorage = append(taskStorage, task)

	fmt.Println("Task Created: ", task.Title, task.DueDate, task.CategoryID)
}

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == AuthenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	c := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: AuthenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)
}
func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, name, password string

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", name, id, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	//findout file exist or not
	var file *os.File

	file, err := os.OpenFile(userStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("can not create or open file: %v\n ", err)

		return
	}

	// save user data in user.txt file
	data := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name, user.Email, user.Password)

	var b = []byte(data)

	numberOfWrittenBytes, wErr := file.Write(b)
	if wErr != nil {
		fmt.Printf("can not write to the file %v\n ", wErr)

		return
	}

	fmt.Println("numberOfWrittenBytes:", numberOfWrittenBytes)

	file.Close()
}
func login() {

	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			AuthenticatedUser = &user
			fmt.Println("you are logged in!")

			break
		}
	}

	if AuthenticatedUser == nil {
		fmt.Println("the email or password is incorrect!")
	}

}


	func loadUserStorageFromFile(){
		file, err := os.Open(userStoragePath)

		if err != nil {
			fmt.Println("can't open the file", err)
		}

		var data = make ([]byte, 1024)
		_ , oErr :=  file.Read(data)
		if oErr != nil {
			fmt.Printf("can't read from the file: %v\n" , oErr)
		}

		var dataStr = string(data)

		userSlice := strings.Split(dataStr , "\n")

		for _ , u := range userSlice {
			if u == " "{
				continue
			}
			//fmt.Println("line of file", index, "user" ,  u)
		    var user = User{}
			userFields := strings.Split(u , ",")
			for _ , field := range userFields {
				//fmt.Println(field)
				values := strings.Split(field , ": ")
				fieldname := strings.ReplaceAll(values[0] , " ", "")
				fieldvalue := values[1]


				switch fieldname{
				case "id":
					id , err := strconv.Atoi(fieldvalue)
					if err != nil {
						fmt.Println("strconv error" , err)

						return
					}
					user.ID = id

				case "name":
					user.Name = fieldvalue
				case "email":
					user.Email = fieldvalue
				case "password":
					user.Password = fieldvalue	 			
				}
			}
				fmt.Printf("user: %+v\n" , user)
		}
	}
