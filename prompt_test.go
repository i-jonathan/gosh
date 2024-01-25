package main

import (
	"testing"
)

func TestGetPrompt(t *testing.T) {
	got, _ := getPrompt()

	if got == "" {
		t.Errorf("Did not get the prompt data")
	}
}
