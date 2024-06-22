package colors

import (
	"os"
	"strings"
)

var shell = os.Args[1]

var plainCode = "\033[0m"
var boldCode = "\033[1m"
var redCode = "\033[31m"
var greenCode = "\033[32m"

func Wrap(colorCode string) string {
	if strings.Contains(shell, "zsh") {
		return "%{" + colorCode + "%}"
	}

	return colorCode
}

func Color(colorCode string, text string, bold bool) string {
	if bold {
		return Wrap(colorCode) + Wrap(boldCode) + text + Wrap(plainCode)
	} else {
		return colorCode + text + plainCode
	}
}

func Plain(text string, bold ...bool) string {
	return Color(plainCode, text, bold[0])
}

func Red(text string, bold ...bool) string {
	return Color(redCode, text, bold[0])
}

func Green(text string, bold ...bool) string {
	return Color(greenCode, text, bold[0])
}

func Yellow(text string, bold ...bool) string {
	return Color("\033[33m", text, bold[0])
}

func Blue(text string, bold ...bool) string {
	return Color("\033[34m", text, bold[0])
}

func Cyan(text string, bold ...bool) string {
	return Color("\033[36m", text, bold[0])
}
