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
		{1, "\x1b[1m"},
		{2, "\x1b[2m"},
		{3, "\x1b[3m"},
		{4, "\x1b[4m"},
		{5, "\x1b[5m"},
		{6, "\x1b[6m"},
	}

	for _, test := range colorTests {
		if result := getColor(test.code); result != test.expected {
			t.Errorf("Output %q is different from expected %q", result, test.expected)
		}
	}
}

func TestFormatText(t *testing.T) {
	textColorTests := []textColorTest{
		{0, "White", "\x1b[0mWhite\x1b[0m"},
		{1, "Red", "\x1b[1mRed\x1b[0m"},
		{2, "Green", "\x1b[2mGreen\x1b[0m"},
		{3, "Yellow", "\x1b[3mYellow\x1b[0m"},
		{4, "Blue", "\x1b[4mBlue\x1b[0m"},
		{5, "Purple", "\x1b[5mPurple\x1b[0m"},
		{6, "Cyan", "\x1b[6mCyan\x1b[0m"},
	}

	for _, test := range textColorTests {
		if output := formatText(test.color, test.text); output != test.expectedText {
			t.Errorf("Output %q is different from expectation %q", output, test.expectedText)
		}
	}
}
