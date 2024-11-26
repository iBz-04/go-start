package greetings

import (
	"errors"
    "fmt"
)


// returns greeting for the named person

func Hello (name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf("Hi, %v. Akwaaba!", name)
	return message, nil
}