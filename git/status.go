package git

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
