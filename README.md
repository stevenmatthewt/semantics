# semantics
A utility written in Golang to automatically handle Semantic Versioning in a CI environment.

[![CircleCI](https://circleci.com/gh/stevenmatthewt/semantics/tree/master.svg?style=shield)](https://circleci.com/gh/stevenmatthewt/semantics/tree/master)

## Install

`go get github.com/stevenmatthewt/semantics`

## Use

```
cd <your github repo>
semantics
```

## What it does

`semantics` is a command-line interface that is intended to be used in a CI environment (or manually, if you're desperate). It automatically handles the process of creating release tags for your project. 

`semantics` relies an a particular commit message structure (completely customizable) to determine what version to create and then automatically creates a release tag and pushes it to GitHub. These release tags comply with [Semantic Version 2.0](semver.org).

At the moment, `semantics` only supports pushing release tags to GitHub, but with a little extra work it can be used to push your release anywhere (npm, dockerhub, quay.io...)

## How it works

In order for `semantics` to work, a particular structure must be adhered to when committing changes to your target repo. By default, the following commit structure will allow for automatic Semantic Versioning off the repo.

 - Commits beginning with `major:` will increase the release tag by one major version.
   - This will reset both the minor, and patch versions to 0.
   - For instance, `1.4.7` would become `2.0.0`
 - Commits beginning with `minor:` will increase the release tag by one minor version.
   - This will reset the patch version to 0.
   - For instance, `1.4.7` would become `1.5.0`
 - Commits beginning with `patch:` will increase the release tag by one patch version.
   - For instance, `1.4.7` would become `1.4.8`
   
## Command Line Arguments

`semantics` is intended to be adapted to whatever workflow suites you best, so there are a number of arguments to the CLI:

- `--major=<pattern>`: use a custom regex pattern to recognize commits that should trigger a major version bump.
- `--minor=<pattern>`: use a custom regex pattern to recognize commits that should trigger a minor version bump.
- `--patch=<pattern>`: use a custom regex pattern to recognize commits that should trigger a patch version bump.

## FAQ

**Q: Do I have to change my git workflow for this to work properly?**

A: Yes. But not much!

In order for `semantics` to work, you need to make sure that any commit that **you want to trigger a version bump** is properly formatted. For instance, if you want to bump one minor version, you'll need a commit with a message like `minor: some stuff I did`. Note that **not all commits must be formatted like this**; you only need to do this when you want to bump version (Merging Pull Requests, for instance)

Additionally, you shouldn't force push to any branch you have set up with `semantics`. If you force push to a release branch, you are overwriting your history, which is going to completely throw `semantics` into a tizzy. Nobody wants that.

**Q: Can I use `semantics` to publish releases to NPM?**

A: Kinda.

At the moment, `semantics` only supports sending tags to GitHub automatically. We're working on NPM support, but it isn't ready at the moment.

However, you can still use `semantics` to help you out with pushing tags to NPM. If you provide the `--output-only` argument, `semantics` will **not** push your tags anywhere, but will just print the new tag to stdout. You can use this in a simple script that can then push that tag to NPM.
