package git

import (
	"errors"
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

	var branchName = string(res)
	return strings.Trim(branchName, "\n"), nil
}

func countCommits(source string, target string) (int, error) {
	var branchRange = source + ".." + target
	var command = exec.Command("git", "rev-list", branchRange)
	var res, err = command.Output()
	if err != nil {
		return 0, err
	}

	var revList = strings.Split(string(res), "\n")
	revList = slices.DeleteFunc(revList, func(e string) bool { return e == "" })

	return len(revList), nil
}

func getCommitCounts(branchName string) (int, int, error) {
	var origin = "origin/" + branchName

	var commitsAhead, aErr = countCommits(origin, "HEAD")
	var commitsBehind, bErr = countCommits("HEAD", origin)
	if aErr != nil || bErr != nil {
		return 0, 0, errors.New("Unable to get commits")
	}

	return commitsAhead, commitsBehind, nil
}
