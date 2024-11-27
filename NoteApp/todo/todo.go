package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("title & content is required")
	}

	return Todo{
		Text: text,
	}, nil
}
func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)

	return nil
}
func (todo Todo) Print() {
	fmt.Println("Text: ", todo.Text)
}
