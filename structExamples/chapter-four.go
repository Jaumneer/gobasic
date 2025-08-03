package main

import (
	"chapter-four/userStruct"
	"fmt"
)

func main() {
	var note userStruct.Note
	err := note.GetNoteContent()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	note.Display()
	err = note.SaveToFile()
	if err != nil {
		fmt.Printf("Error saving note: %v\n", err)
		return
	}
	fmt.Println("Note saved successfully!")
}
