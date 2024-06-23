package git

type GitStatus int

const (
	Equal GitStatus = iota
	Ahead
	Behind
	Diverged
)

type LocalGitStatus int

const (
	None LocalGitStatus = iota
	Staged
	Unstaged
	Both
)

const GitIconEqual string = "≡"
const GitIconAhead string = "↑"
const GitIconBehind string = "↓"
const GitIconDiverged string = "↕"
