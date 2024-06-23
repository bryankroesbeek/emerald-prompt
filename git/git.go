package git

import (
	"strings"
	"emerald/colors"
)

type SyncStatus struct {
	status GitStatus
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

func GetBranch() string {
	var branch, err = getBranchName()
	if err != nil || branch == "" {
		return ""
	}

	return " (" + strings.TrimSpace(branch) + ")"
}
