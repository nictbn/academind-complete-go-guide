package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title string
	content string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshall(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(title string, content string) (Note, error) {
	if title == "" || content == ""{
		return Note{}, errors.New("Invalid input!")
	}

	return Note {
		title: title,
		content: content,
		CreatedAt: time.Now(),
	}, nil
}