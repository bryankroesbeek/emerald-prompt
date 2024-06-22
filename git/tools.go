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

func GetStatus() string {
	var cmd = exec.Command("git", "status", "--porcelain")
	var res, err = cmd.Output()
	if err != nil {
		return ""
	}

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
