package greetings

import (
	"errors"
    "fmt"
	"math/rand"
)


// returns greeting for the named person

func Hello (name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty")
	}

	// create random message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying a random index
	return formats[rand.Intn(len(formats))]
}	