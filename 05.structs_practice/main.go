package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo: ")

	noteData, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	todoData, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	noteData.Display()
	todoData.Display()

	err = saveData(noteData)
	if err != nil {
		return
	}

	err = saveData(todoData)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Saving the data failed")
		return err
	}

	fmt.Println("Saving the data succeded")
	return nil
}

func getNoteData() (string, string) {
	return getUserInput("Note Title: "), getUserInput("Note Content: ")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
