package git

import (
	"os/exec"
	"strings"
)

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
	var cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var res, err = cmd.Output()
	if err != nil {
		return ""
	}
	var branch = string(res)

	if branch == "" {
		return ""
	}

	return " (" + strings.TrimSpace(branch) + ")"
}
