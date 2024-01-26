package main

import (
	"testing"
)

type colorTest struct {
	code     int
	expected string
}

type textColorTest struct {
	color              int
	text, expectedText string
}

func TestGetColor(t *testing.T) {
	colorTests := []colorTest{
		{0, "\x1b[0m"},
		{31, "\x1b[31m"},
		{32, "\x1b[32m"},
		{33, "\x1b[33m"},
		{34, "\x1b[34m"},
		{35, "\x1b[35m"},
		{36, "\x1b[36m"},
	}

	for _, test := range colorTests {
		if result := getTextFormat(test.code); result != test.expected {
			t.Errorf("Output %q is different from expected %q", result, test.expected)
		}
	}
}

func TestFormatText(t *testing.T) {
	textColorTests := []textColorTest{
		{0, "Reset", "\x1b[0mReset\x1b[0m"},
		{31, "Red", "\x1b[31mRed\x1b[0m"},
		{32, "Green", "\x1b[32mGreen\x1b[0m"},
		{33, "Yellow", "\x1b[33mYellow\x1b[0m"},
		{34, "Blue", "\x1b[34mBlue\x1b[0m"},
		{35, "Purple", "\x1b[35mPurple\x1b[0m"},
		{36, "Cyan", "\x1b[36mCyan\x1b[0m"},
	}

	for _, test := range textColorTests {
		if output := formatText(test.color, test.text); output != test.expectedText {
			t.Errorf("Output %q is different from expectation %q", output, test.expectedText)
		}
	}
}
