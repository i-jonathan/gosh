package main

import (
	"fmt"
	"os"
)

func getPrompt() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n> ", pwd), nil
}
