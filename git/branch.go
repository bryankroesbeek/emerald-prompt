package git

import (
	"errors"
	"fmt"
	"os/exec"
	"slices"
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

func countCommits(source string, target string) (int, error) {
	var aheadCmd = exec.Command("git", "rev-list", fmt.Sprint(source, "..", target))
	var res, err = aheadCmd.Output()
	if err != nil {
		return 0, err
	}

	var revList = strings.Split(string(res), "\n")
	revList = slices.DeleteFunc(revList, func(e string) bool { return e == "" })

	return len(revList), nil
}

func getCommitCounts(branchName string) (int, int, error) {
	var origin = "origin/" + branchName

	var commitsAhead, aErr = countCommits("HEAD", origin)
	var commitsBehind, bErr = countCommits(origin, "HEAD")
	if aErr != nil || bErr != nil {
		return 0, 0, errors.New("Unable to get commits")
	}

	return commitsBehind, commitsAhead, nil
}
