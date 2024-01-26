package main

import (
	"testing"
)

type formatTest struct {
	code     int
	expected string
}

type textFormatTest struct {
	color              int
	text, expectedText string
}

func TestGetTextFormat(t *testing.T) {
	colorTests := []formatTest{
		{RESET, "\x1b[0m"},
		{BOLD, "\x1b[1m"},
		{ITALIC, "\x1b[3m"},
		{UNDERLINE, "\x1b[4m"},
		{INVERT, "\x1b[7m"},
		{CROSSED, "\x1b[9m"},
		{RED, "\x1b[31m"},
		{GREEN, "\x1b[32m"},
		{YELLOW, "\x1b[33m"},
		{BLUE, "\x1b[34m"},
		{PURPLE, "\x1b[35m"},
		{CYAN, "\x1b[36m"},
	}

	for _, test := range colorTests {
		if result := getTextFormat(test.code); result != test.expected {
			t.Errorf("Output %q is different from expected %q", result, test.expected)
		}
	}
}

func TestFormatText(t *testing.T) {
	textColorTests := []textFormatTest{
		{RESET, "Reset", "\x1b[0mReset\x1b[0m"},
		{BOLD, "Bold", "\x1b[1mBold\x1b[0m"},
		{ITALIC, "Italics", "\x1b[3mItalics\x1b[0m"},
		{UNDERLINE, "Underline", "\x1b[4mUnderline\x1b[0m"},
		{INVERT, "Invert", "\x1b[7mInvert\x1b[0m"},
		{CROSSED, "Crossed", "\x1b[9mCrossed\x1b[0m"},
		{RED, "RedText", "\x1b[31mRedText\x1b[0m"},
		{GREEN, "GreenText", "\x1b[32mGreenText\x1b[0m"},
		{YELLOW, "YellowText", "\x1b[33mYellowText\x1b[0m"},
		{BLUE, "BlueText", "\x1b[34mBlueText\x1b[0m"},
		{PURPLE, "PurpleText", "\x1b[35mPurpleText\x1b[0m"},
		{CYAN, "CyanText", "\x1b[36mCyanText\x1b[0m"},
		{WHITE, "WhiteText", "\x1b[37mWhiteText\x1b[0m"},
	}

	for _, test := range textColorTests {
		output := formatText(test.color, test.text)
		t.Log(output)
		if output != test.expectedText {
			t.Errorf("Output %q is different from expectation %q", output, test.expectedText)
		}
	}
}
