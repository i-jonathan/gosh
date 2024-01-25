package main

import (
	"fmt"
)

const escape = "\x1b"

const (
	NONE = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
)

func getColor(code int) string {
	return fmt.Sprintf("%s[%dm", escape, code)
}

func formatText(color int, text string) string {
	return getColor(color) + text + getColor(NONE)
}
