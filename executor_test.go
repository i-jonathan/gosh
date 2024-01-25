package main

import (
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	got := executeCommand("ls")

	if got != nil {
		t.Error("Most Basic command failed")
	}
}
