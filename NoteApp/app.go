package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note"
	"example.com/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteInput()
	userNote, error := note.New(title, content)

	if error != nil {
		fmt.Print(error)
		return
	}

	todoText := getInput("Todo: ")
	todo, error := todo.New(todoText)

	if error != nil {
		fmt.Print(error)
		return
	}

	saveData(userNote)
	saveData(todo)

	userNote.Print()
	todo.Print()
}
func saveData(data saver) error {
	error := data.Save()
	if error != nil {
		fmt.Print(error)
		return error
	}
	return nil
}

func getNoteInput() (string, string) {
	title := getInput("Enter title: ")
	content := getInput("Enter content: ")

	return title, content
}

func getInput(prompt string) string {
	value := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	value.Scan()

	return value.Text()
}
