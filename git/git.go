package git

import (
	"strings"
)


func GetBranch() string {
	var branch, err = getBranchName()
	if err != nil || branch == "" {
		return ""
	}

	return " (" + strings.TrimSpace(branch) + ")"
}
