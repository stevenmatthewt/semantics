package main

import (
	"flag"
	"fmt"

	"github.com/stevenmatthewt/semantics/bump"
	"github.com/stevenmatthewt/semantics/git"
	"github.com/stevenmatthewt/semantics/output"
)

// CLIFlags stores all flags that can be passed through the CLI
type CLIFlags struct {
	major     string
	minor     string
	patch     string
	outputTag bool
	dry       bool
}

func main() {
	flags := getFlags()
	if flags.outputTag {
		output.PrintToStdout = false
	}
	output.Stdout(fmt.Sprintf("major: %v, minor: %v, patch: %v\n", flags.major, flags.minor, flags.patch))

	tag, err := git.GetLatestTag()
	if err != nil {
		output.Fatal(err)
	}
	output.Stdout(fmt.Sprintf("old tag: %v.%v.%v\n", tag.Major, tag.Minor, tag.Patch))
	commits := git.GetCommitsSinceTag(tag)

	bumpMap, err := bump.MapFromStrings(flags.major, flags.minor, flags.patch)
	if err != nil {
		output.Fatal(fmt.Sprintf("One of the regexes provided did not compile: %v", err))
	}
	bumps := commits.ScanForBumps(bumpMap)
	if len(bumps) == 0 {
		output.Stdout("No updates to version. Aborting.")
		return
	}

	for _, b := range bumps {
		tag = b.Bump(tag)
	}
	output.Stdout(fmt.Sprintf("new tag: %s\n", tag.String()))
	if flags.outputTag {
		output.StdoutForce(tag.String())
	}

	if flags.dry == false {
		err = git.PushTag(tag)
		if err != nil {
			output.Fatal(err)
		}
	}
}

func getFlags() (flags CLIFlags) {
	flag.StringVar(&flags.major, "major", "^major:.*", "Commit tag regex that indicates a Major bump should be performed.")
	flag.StringVar(&flags.minor, "minor", "^minor:.*", "Commit tag regex that indicates a Minor bump should be performed.")
	flag.StringVar(&flags.patch, "patch", "^patch:.*", "Commit tag regex that indicates a Patch bump should be performed.")
	flag.BoolVar(&flags.outputTag, "output-tag", false, "Print only the new tag to stdout. Usually combined with dry-run. Default: false")
	flag.BoolVar(&flags.dry, "dry-run", false, "Don't create new tag, or push to github. Default: false")
	flag.Parse()

	return flags
}
