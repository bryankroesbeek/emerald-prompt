package git

import (
	"os/exec"
	"strings"
)

// BUG: When a rename is in the status, it will show up as both a tracked and untracked change
func getLocalCounts() (int, int, error) {
	var cmd = exec.Command("git", "status", "-z")
	var res, err = cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var changesStr = string(res)
	var changes StringSlice = strings.Split(changesStr, "\000")
	changes = changes.filterEmpty()
	var countI, countW = countChanges(changes)
	return countI, countW, nil
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
