package git

import (
	"fmt"
	"os/exec"
	"strings"
)

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
