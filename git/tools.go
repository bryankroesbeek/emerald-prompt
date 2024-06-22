package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func getBranchName() (string, error) {
	var cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var res, err = cmd.Output()
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func countChanges(changes []string) (int, int) {
	var index int = 0
	var workingTree int = 0

	for _, change := range changes {
		if change == "" {
			continue
		}
		var cIndex = string(change[0])
		var tIndex = string(change[1])

		if string(change[0:2]) == "??" {
			workingTree = workingTree + 1
			continue
		}

		if cIndex != " " {
			index = index + 1
		}
		if tIndex != " " {
			workingTree = workingTree + 1
		}
	}

	return index, workingTree
}

func GetStatus() string {
	var cmd = exec.Command("git", "status", "-z")
	var res, err = cmd.Output()
	if err != nil {
		return ""
	}

	var changesStr = string(res)
	var changes = strings.Split(changesStr, "\000")
	var countI, countW = countChanges(changes)
	fmt.Println(countI, countW)

	if len(res) == 0 {
		return ":)"
	}

	return ":("
}

func GetBranch() string {
	var branch, err = getBranchName()
	if err != nil || branch == "" {
		return ""
	}

	return " (" + strings.TrimSpace(branch) + ")"
}
