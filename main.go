package main

import (
	"emerald/colors"
	"emerald/git"
	"fmt"
	"os"
	"strings"
)

var green = colors.Green
var plain = colors.Plain
var cyan = colors.Cyan

var user = os.Getenv("USER")
var host, _ = os.Hostname()
var dir, _ = os.Getwd()
var gitBranch = git.GetBranch()
var gitStatus = git.GetStatus()

func getDir() string {
	var dir, _ = os.Getwd()
	var homeDir = os.Getenv("HOME")

	if dir == "/" {
		return dir
	}

	var path = strings.Replace(dir, homeDir, "~", 1)
	var parts = strings.Split(path, "/")

	return parts[len(parts)-1]
}

func main() {
	var prompt = fmt.Sprint(
		green("["+user+"@"+host+" ", true),
		plain(getDir(), true),
		green("]", true),
		cyan(gitBranch, true),
		green("$ ", true),
	)

	fmt.Println(prompt)
}
