package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/stevenmatthewt/semantics/bump"
	"github.com/stevenmatthewt/semantics/git"
)

// CLIFlags stores all flags that can be passed through the CLI
type CLIFlags struct {
	major string
	minor string
	patch string
}

func main() {
	flags := getFlags()
	fmt.Printf("major: %v, minor: %v, patch: %v\n", flags.major, flags.minor, flags.patch)
	tag, err := git.GetLatestTag()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("old tag: %v.%v.%v\n", tag.Major, tag.Minor, tag.Patch)
	commits := git.GetCommitsSinceTag(tag)

	bumpMap, err := bump.MapFromStrings(flags.major, flags.minor, flags.patch)
	if err != nil {
		log.Fatalf("One of the regexes provided did not compile: %v", err)
	}
	bumps := commits.ScanForBumps(bumpMap)
	if len(bumps) == 0 {
		log.Print("No updates to version. Aborting.")
		return
	}

	for _, b := range bumps {
		tag = b.Bump(tag)
	}
	fmt.Printf("new tag: %v.%v.%v\n", tag.Major, tag.Minor, tag.Patch)

	err = git.PushTag(tag)
	if err != nil {
		log.Fatal(err)
	}
}

func getFlags() (flags CLIFlags) {
	flag.StringVar(&flags.major, "major", "^major:.*", "Commit tag regex that indicates a Major bump should be performed.")
	flag.StringVar(&flags.minor, "minor", "^minor:.*", "Commit tag regex that indicates a Minor bump should be performed.")
	flag.StringVar(&flags.patch, "patch", "^patch:.*", "Commit tag regex that indicates a Patch bump should be performed.")
	flag.Parse()

	return flags
}
