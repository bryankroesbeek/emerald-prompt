package main

import (
	"emerald/git"
	"fmt"
	"os"
	"strings"
)

var shell = os.Args[1]
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
		Green("["+user+"@"+host+" ", true),
		Plain(getDir(), true),
		Green("]", true),
		Cyan(gitBranch, true),
		Green("$ ", true),
	)

	fmt.Println(prompt)
}
