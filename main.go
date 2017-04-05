package main

import (
	"flag"
	"fmt"

	"github.com/cbdr/semantics/bump"
	"github.com/cbdr/semantics/git"
)

type CLIFlags struct {
	major string
	minor string
	patch string
}

func main() {
	flags := getFlags()
	fmt.Printf("major: %v, minor %v, patch %v\n", flags.major, flags.minor, flags.patch)

	tag := git.GetLatestTag()
	fmt.Printf("%+v\n", tag)
	commits := git.GetCommitsSinceTag(tag)
	fmt.Printf("%+v\n", commits)

	bumpMap := getBumpMap(flags)
	major, minor, patch := commits.ScanForBumps(bumpMap)
	fmt.Printf("major: %v, minor %v, patch %v\n", major, minor, patch)
}

func getFlags() (flags CLIFlags) {
	flag.StringVar(&flags.major, "major", "major", "Commit tag that indicates a Major bump should be performed.")
	flag.StringVar(&flags.minor, "minor", "minor", "Commit tag that indicates a Minor bump should be performed.")
	flag.StringVar(&flags.patch, "patch", "patch", "Commit tag that indicates a Patch bump should be performed.")
	flag.Parse()

	return flags
}

func getBumpMap(flags CLIFlags) bump.Map {
	return bump.Map{
		Major: flags.major,
		Minor: flags.minor,
		Patch: flags.patch,
	}
}
