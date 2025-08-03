package userStruct

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title   string
	Content string
	Time    time.Time
}

func (u *Note) GetNoteContent() error {
	fmt.Printf("Please Enter The Title:")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')

	u.Title = strings.TrimSpace(title)
	if u.Title == "" {
		return errors.New("title cannot be empty")
	}

	fmt.Printf("\nPlease Enter The Content:")
	content, _ := reader.ReadString('\n')

	u.Content = strings.TrimSpace(content)
	if u.Content == "" {
		return errors.New("content cannot be empty")
	}
	u.Time = time.Now()
	return nil
}

func (u *Note) Display() {
	fmt.Printf("Title: %s\n", u.Title)
	fmt.Printf("Content: %s\n", u.Content)
	fmt.Printf("Time: %s\n", u.Time.Format(time.RFC1123))
}

func (u *Note) SaveToFile() error {
	fileName := strings.ReplaceAll(u.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	data, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("error marshaling note to JSON: %w", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing note to file: %w", err)
	}

	return nil
}
