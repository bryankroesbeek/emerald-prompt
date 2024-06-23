package git

import "slices"

type StringSlice []string

func (s StringSlice) filterEmpty() []string {
	return slices.DeleteFunc(s, func(in string) bool { return in == "" })
}
