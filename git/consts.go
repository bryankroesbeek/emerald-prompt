package git

type GitStatus int

const (
	Equal GitStatus = iota
	Ahead
	Behind
	Diverged
)

const GitIconEqual string = "≡"
const GitIconAhead string = "↑"
const GitIconBehind string = "↓"
const GitIconDiverged string = "↕"
