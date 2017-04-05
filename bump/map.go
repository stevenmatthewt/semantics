package bump

import "regexp"

// Map represents a mapping of commit message regexes
// to what they should increment. I.e. Major, Minor, Patch
type Map struct {
	Major *regexp.Regexp
	Minor *regexp.Regexp
	Patch *regexp.Regexp
}

// MapFromStrings takes three strings that describe regexes to match
// Semantic Versions on.
func MapFromStrings(major, minor, patch string) (m Map, err error) {
	if m.Major, err = regexp.Compile(major); err != nil {
		return Map{}, err
	}
	if m.Minor, err = regexp.Compile(minor); err != nil {
		return Map{}, err
	}
	if m.Patch, err = regexp.Compile(patch); err != nil {
		return Map{}, err
	}

	return m, nil
}
