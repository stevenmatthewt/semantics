package main

import (
	"flag"
	"fmt"

	"github.com/cbdr/semantics/git"
)

type flags struct {
	major *string
	minor *string
	patch *string
}

func main() {
	flags := getFlags()
	fmt.Printf("major: %v, minor %v, patch %v\n", *flags.major, *flags.minor, *flags.patch)

	tag := git.GetLatestTag()
	fmt.Printf("%+v\n", tag)
	commits := git.GetCommitsSinceTag(tag)
	fmt.Printf("%+v\n", commits)
}

func getFlags() flags {
	var f flags
	f.major = flag.String("major", "major", "Commit tag that indicates a Major bump should be performed.")
	f.minor = flag.String("minor", "minor", "Commit tag that indicates a Minor bump should be performed.")
	f.patch = flag.String("patch", "patch", "Commit tag that indicates a Patch bump should be performed.")
	flag.Parse()

	return f
}
