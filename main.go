package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cbdr/semantics/bump"
	"github.com/cbdr/semantics/git"
)

type CLIFlags struct {
	major string
	minor string
	patch string
}

type tag string

func main() {
	flags := getFlags()
	fmt.Printf("major: %v, minor: %v, patch: %v\n", flags.major, flags.minor, flags.patch)
	tag, err := git.GetLatestTag()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("old tag: %v.%v.%v\n", tag.Major, tag.Minor, tag.Patch)
	commits := git.GetCommitsSinceTag(tag)

	bumpMap := getBumpMap(flags)
	bumps := commits.ScanForBumps(bumpMap)

	for _, b := range bumps {
		tag = b.Bump(tag)
	}
	fmt.Printf("new tag: %v.%v.%v\n", tag.Major, tag.Minor, tag.Patch)
}

func getFlags() (flags CLIFlags) {
	flag.StringVar(&flags.major, "major", "major", "Commit tag that indicates a Major bump should be performed.")
	flag.StringVar(&flags.minor, "minor", "minor", "Commit tag that indicates a Minor bump should be performed.")
	flag.StringVar(&flags.patch, "patch", "patch", "Commit tag that indicates a Patch bump should be performed.")
	flag.Parse()

	validateTag(flags.major)
	validateTag(flags.minor)
	validateTag(flags.patch)

	return flags
}

func getBumpMap(flags CLIFlags) bump.Map {
	return bump.Map{
		Major: flags.major,
		Minor: flags.minor,
		Patch: flags.patch,
	}
}

func validateTag(t string) {
	if t == "" {
		log.Fatal("Not a valid tag: \"\"")
	}
}
