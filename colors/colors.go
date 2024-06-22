package colors

import (
	"os"
	"strings"
)

var shell = os.Args[1]

const plainCode = "\033[0m"
const boldCode = "\033[1m"
const redCode = "\033[31m"
const greenCode = "\033[32m"
const yellowCode = "\033[33m"
const blueCode = "\033[34m"
const purpleCode = "\033[35m"
const cyanCode = "\033[36m"
const whiteCode = "\033[37m"

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
	return Color(yellowCode, text, bold[0])
}

func Blue(text string, bold ...bool) string {
	return Color(blueCode, text, bold[0])
}

func Purple(text string, bold ...bool) string {
	return Color(purpleCode, text, bold[0])
}

func Cyan(text string, bold ...bool) string {
	return Color(cyanCode, text, bold[0])
}

func White(text string, bold ...bool) string {
	return Color(whiteCode, text, bold[0])
}
