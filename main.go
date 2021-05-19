package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	var task []Task
	menu := [][]string{
		[]string{"1", "Add"},
		[]string{"2", "Show"},
		[]string{"3", "Delete"},
	}
	tableMenu := tablewriter.NewWriter(os.Stdout)
	tableMenu.SetHeader([]string{"Menu", "Action"})
	for _, v := range menu {
		tableMenu.Append(v)
	}
	tableMenu.Render()

	data := [][]string{}

	scanner := bufio.NewScanner(os.Stdin)
	var userInput string
	for userInput != "exits" { // break the loop if text == "exits"
		fmt.Print(">>> ")
		scanner.Scan()
		userInput = scanner.Text()

		if userInput == "1" {
			var userInputTaskId string
			fmt.Print(">>> Id: ")
			scanner.Scan()
			userInputTaskId = scanner.Text()

			var userInputTaskTitle string
			fmt.Print(">>> Title: ")
			scanner.Scan()
			userInputTaskTitle = scanner.Text()

			var userInputTaskContent string
			fmt.Print(">>> Content: ")
			scanner.Scan()
			userInputTaskContent = scanner.Text()
			task = append(task, Task{ID: userInputTaskId, Title: userInputTaskTitle, Content: userInputTaskContent})
		}

		if userInput == "2" {
			data = data[:0][:0]
			for _, v := range task {
				data = append(data, []string{v.ID, v.Title, v.Content})

			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Title", "Content"})
			for _, v := range data {

				table.Append(v)
			}
			table.Render()
		}

		if userInput == "3" {
			var userinputTaskId string
			fmt.Print(">>> Id: ")
			scanner.Scan()
			userinputTaskId = scanner.Text()

			for i := 0; i < len(task); i++ {
				if task[i].ID == userinputTaskId {
					task = append(task[:i], task[i+1:]...)
				}
			}

			data = data[:0][:0]
			for _, v := range task {
				data = append(data, []string{v.ID, v.Title, v.Content})
			}
		}
	}
}
