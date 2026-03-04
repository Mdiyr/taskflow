package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct{
	ID int
	Name string
	Email string
	Password string
}

type Task struct{
	ID int
	UserID int
	Title string
	DueDate	string
	Category string
	IsDone bool
}

var userStorage []User
var AuthenticatedUser *User

var taskStorage []Task

func (u User) print() {
	fmt.Println("User:", u.ID, u.Email, u.Name)
}


func main() {
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()



	for{

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


		task := Task {
			ID: len(taskStorage)+1,
			UserID: AuthenticatedUser.ID,
			Title: title,
			DueDate: duedate,
			Category: category,
			IsDone: false,
		}
		taskStorage = append(taskStorage, task)

		fmt.Println("Task Created: ", task.Title, task.DueDate, task.Category)
}

func listTask() {
	for _ ,task := range taskStorage {
		if task.UserID == AuthenticatedUser.ID{
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

		fmt.Println("category:", title , color)
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
			ID: len(userStorage) + 1,
			Name: name,
			Email: email,
			Password: password,
		}

		userStorage = append(userStorage, user)
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

/* how to get info from cli step by step
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*command = scanner.Text()
	*/

