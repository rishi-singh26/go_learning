package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	noteData, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	noteData.Display()
	err = noteData.Save()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeded")
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
