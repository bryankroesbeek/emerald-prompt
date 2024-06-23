package git

import (
	"emerald/colors"
	"fmt"
)

type SyncStatus struct {
	status GitStatus
	icon   string
	color  func(string, bool) string
}

type LocalStatus struct {
	status LocalGitStatus
	icon   string
	color  func(string, bool) string
}

func getSyncStatus(ahead int, behind int) SyncStatus {
	if ahead > 0 && behind > 0 {
		return SyncStatus{Diverged, GitIconDiverged, colors.Yellow}
	}

	if ahead > 0 {
		return SyncStatus{Ahead, GitIconAhead, colors.Green}
	}

	if behind > 0 {
		return SyncStatus{Behind, GitIconBehind, colors.Red}
	}

	return SyncStatus{Equal, GitIconEqual, colors.Cyan}
}

func getLocalStatus(index int, workingTree int) LocalStatus {
	if index > 0 && workingTree > 0 {
		return LocalStatus{Both, "!", colors.Yellow}
	}

	if index > 0 {
		return LocalStatus{Staged, "!", colors.Green}
	}

	if workingTree > 0 {
		return LocalStatus{Unstaged, "!", colors.Red}
	}

	return LocalStatus{None, "", colors.Cyan}
}

func GetStatus() string {
	var branchName, err = getBranchName()
	if err != nil || branchName == "" {
		return ""
	}

	var commitsAhead, commitsBehind, branchErr = getCommitCounts(branchName)
	if branchErr != nil {
		return ""
	}

	var gitSync = getSyncStatus(commitsAhead, commitsBehind)
	var start = gitSync.color(" (", true)
	var branch = gitSync.color(branchName+" ", true)
	var status = gitSync.color(gitSync.icon, true)
	var end = gitSync.color(")", true)

	if gitSync.status == Equal {
		var index, workingTree, err = getLocalCounts()
		if err != nil {
			return ""
		}
		var gitLocal = getLocalStatus(index, workingTree)
		if gitLocal.status != None {
			status = gitLocal.color(gitLocal.icon, true)
		}
	}

	return fmt.Sprint(
		start,
		branch,
		status,
		end,
	)
}
