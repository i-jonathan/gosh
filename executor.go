package main

import (
	"os"
	"os/exec"
	"strings"
)

func executeCommand(command string) error {
	command = strings.TrimSuffix(command, "\n")
	args := strings.Split(command, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
		return nil
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
