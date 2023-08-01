/*
GoTaskManager.go
Author: Sergio Lima
Creation date: 29 July 2023

GoTaskManager is a command-line utility written in Go.
It provides a task management system where users can create, update, delete, and list tasks.
Each task consists of a description, a description, a status, a due date, and a completion date.

Tasks are stored in a local JSON file descriptiond 'tasks.json'. The file is loaded when the program starts,
and any changes made during the program's execution are saved back to the file when the program is quit.
This allows for persistent storage of tasks between different executions of the application.

The user interacts with GoTaskManager through a menu system:

MENU
============
C - CREATE
U - UPDATE
D - DELETE
L - LIST
O - COMPLETE
Q - QUIT

For example:
Choosing 'C' prompts the user to enter the description of the task and its due date, creating a new task.
Choosing 'U' prompts the user to enter the ID of the task to update and the new description for that task.
Choosing 'D' prompts the user to enter the ID of the task to delete.
Choosing 'L' lists all the existing tasks.
Choosing 'O' Complete task.
Choosing 'Q' saves the current tasks to the JSON file and quits the program.

To compile this program for Linux, use the following command:
    go build -o GoTaskManager

To compile this program for Windows, use the following command:
    GOOS=windows GOARCH=amd64 go build -o GoTaskManager.exe
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

type Task struct {
	ID            int
	Description   string
	Status        string
	DueDate       time.Time
	CompletedDate time.Time
}

var tasks = make(map[int]Task)
var id = 1
var storagePath = "tasks.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	loadTasks()

	for {
		printWelcomeScreen()
		printMenu()
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		line := scanner.Text()
		tokens := strings.Split(strings.ToUpper(line), " ")
		switch tokens[0] {
		case "Q":
			saveTasks()
			return
		case "C":
			handleCreateTask(scanner)
		case "U":
			handleUpdateTask(scanner)
		case "D":
			handleDeleteTask(scanner)
		case "L":
			listTasks()
			waitForUserInput()
		case "O":
			handleCompleteTask(scanner)
		default:
			fmt.Println("Unknown command:", tokens[0])
		}
	}
}

// Returns task and a bool indicating if it exists
func getTask(scanner *bufio.Scanner, operation string) (int, bool) {
	fmt.Printf("Enter the task ID to %s: ", operation)
	scanner.Scan()
	var id int
	fmt.Sscan(scanner.Text(), &id)
	_, found := tasks[id]
	if !found {
		fmt.Printf("No task found with ID %d\n", id)
	}
	return id, found
}

func handleUpdateTask(scanner *bufio.Scanner) {
	idToUpdate, found := getTask(scanner, "update")
	if !found {
		waitForUserInput()
		return
	}
	fmt.Println("Enter the new description of the task: ")
	scanner.Scan()
	newDescription := scanner.Text()
	task := tasks[idToUpdate]
	task.Description = newDescription
	tasks[idToUpdate] = task
	fmt.Println("Task", idToUpdate, "updated.")
	waitForUserInput()
}

func handleDeleteTask(scanner *bufio.Scanner) {
	idToDelete, found := getTask(scanner, "delete")
	if !found {
		waitForUserInput()
		return
	}
	delete(tasks, idToDelete)
	fmt.Println("Task", idToDelete, "deleted.")
	waitForUserInput()
}

func handleCompleteTask(scanner *bufio.Scanner) {
	idToComplete, found := getTask(scanner, "complete")
	if !found {
		waitForUserInput()
		return
	}
	task := tasks[idToComplete]
	task.Status = "Complete"
	task.CompletedDate = time.Now()
	tasks[idToComplete] = task
	fmt.Println("Task", idToComplete, "marked as complete.")
	waitForUserInput()
}

func handleCreateTask(scanner *bufio.Scanner) {
	fmt.Println("Enter the description of the task: ")
	scanner.Scan()
	description := scanner.Text()
	fmt.Println("Enter the due date (YYYY-MM-DD): ")
	scanner.Scan()
	dueDate, _ := time.Parse("2006-01-02", scanner.Text())
	createTask(description, dueDate)
	waitForUserInput()
}

func waitForUserInput() {
	fmt.Println("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printWelcomeScreen() {
	clearScreen()
	myFigure := figure.NewFigure("GoTaskManager", "", true)
	myFigure.Print()
	fmt.Println("Today's date:", time.Now().Format("2006-01-02"))
	fmt.Println("Current time:", time.Now().Format("15:04:05"))
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println()
}

func printMenu() {
	fmt.Println("MENU")
	fmt.Println("=============")
	fmt.Println("C - CREATE")
	fmt.Println("U - UPDATE")
	fmt.Println("D - DELETE")
	fmt.Println("L - LIST")
	fmt.Println("O - COMPLETE")
	fmt.Println("Q - QUIT")
	fmt.Println()
}

func createTask(description string, dueDate time.Time) {
	task := Task{ID: id, Description: description, DueDate: dueDate, Status: "Pending"}
	tasks[id] = task
	fmt.Println("Created task", id)
	id++
}

func listTasks() {
	printWelcomeScreen()
	// Create a slice of the keys (task IDs)
	taskIDs := make([]int, 0, len(tasks))
	for id := range tasks {
		taskIDs = append(taskIDs, id)
	}

	// Sort the slice in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(taskIDs)))

	fmt.Printf("%4s %-50s %-10s %-20s %-20s\n", "Id", "Description", "Status", "Due Date", "Completed Date")
	fmt.Println(strings.Repeat("=", 140))
	for _, task := range taskIDs {
		completed := tasks[task].CompletedDate.Format("2006-01-02")
		if tasks[task].Status != "Complete" {
			completed = "Not completed yet"
		}
		fmt.Printf("%4d %-50s %-10s %-20s %-20s\n", tasks[task].ID, tasks[task].Description, tasks[task].Status, tasks[task].DueDate.Format("2006-01-02"), completed)
	}
}

func loadTasks() {
	byteValue, err := ioutil.ReadFile(storagePath)
	if err != nil {
		return
	}
	json.Unmarshal(byteValue, &tasks)

	if len(tasks) > 0 {
		for _, task := range tasks {
			if task.ID >= id {
				id = task.ID + 1
			}
		}
	}
}

func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", " ")
	_ = ioutil.WriteFile(storagePath, file, 0644)
}
