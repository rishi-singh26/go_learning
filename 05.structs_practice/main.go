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

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething("Hello")
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

	err = outPutData(noteData)
	if err != nil {
		return
	}

	err = outPutData(todoData)
	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	typedVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", typedVal)
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float 64:", value)
	// default:
	// 	fmt.Println(value)
	// }
}

func outPutData(data outputtable) error {
	data.Display()
	return saveData(data)
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
