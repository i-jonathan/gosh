package main

import (
	"fmt"
)

const escape = "\x1b"

const (
	NONE = iota
	BOLD
	_
	ITALIC
	UNDERLINE
	_
	_
	INVERT
	RED = iota + 23
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
)

func getTextFormat(code int) string {
	return fmt.Sprintf("%s[%dm", escape, code)
}

func formatText(color int, text string) string {
	return getTextFormat(color) + text + getTextFormat(NONE)
}
