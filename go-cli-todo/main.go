package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string // In-memory task storage

func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Added task:", task)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}
	fmt.Println("\n--- TODO ---")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
	fmt.Println("------------")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Go CLI Todo App")
	fmt.Println("----------------------")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command:", err)
			continue
		}

		command := strings.TrimSpace(input)
		parts := strings.Fields(command) // Split command and arguments

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) > 1 {
				taskDescription := strings.Join(parts[1:], " ")
				addTask(taskDescription)
			} else {
				fmt.Println("Usage: add <task description>")
			}
		case "list":
			listTasks()
		case "exit":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Unknown command. Available: add, list, exit")
		}
	}
}
