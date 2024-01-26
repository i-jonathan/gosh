package main

import (
	"fmt"
)

const escape = "\x1b"

const (
	RESET = iota
	BOLD
	_
	ITALIC
	UNDERLINE
	_
	_
	INVERT
	_
	CROSSED
)

const (
	RED = iota + 31
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
	WHITE
)

const (
	HRED = iota + 90
	HGREEN
	HYELLOW
	HBLUE
	HPURPLE
	HCYAN
	HWHITE
)

const (
	BGRED = iota + 41
	BGGREEN
	BGYELLOW
	BGBLUE
	BGPURPLE
	BGCYAN
	BGWHITE
)

const (
	BGHRED = iota + 100
	BGHGREEN
	BGHYELLOW
	BGHBLUE
	BGHPURPLE
	BGHCYAN
	BGHWHITE
)

func getTextFormat(code int) string {
	return fmt.Sprintf("%s[%dm", escape, code)
}

func formatText(color int, text string) string {
	return getTextFormat(color) + text + getTextFormat(RESET)
}
