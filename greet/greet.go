package greet

import (
	"errors"
	"fmt"
)

func Greeting(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	msg := fmt.Sprintf("Hi %v Good One", name)
	return msg, nil
}
