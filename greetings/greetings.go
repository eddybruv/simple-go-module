package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//if no name was provided
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting
// message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting
func randomFormat() string {
	// a slice of formats
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you %v!",
		"What's up %v?",
	}
	length := len(formats)
	return formats[rand.Intn(length)]
}
