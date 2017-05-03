#!/bin/bash -e

packages=$(go list ./... | grep -v vendor | xargs)

usage() {
    cat <<EOF
commands:
    test - run linter, vetter, and tests
EOF
}

go_fmt() {
    echo 'running go fmt on all packages...'
    fmt=$(go fmt $packages 2>&1)
    if [[ $fmt ]]; then
    echo $fmt
    exit 1
    fi
}

go_vet() {
    echo 'running go vet on all packages...'
    go vet $packages
}

go_test() {
    echo 'running go test on all packages...'
    for pkg in $packages
    do
        pkg_name=$(echo "$pkg" | tr / -)
        pkg_cov="${pkg_name}.cov"
        go test -v -race -covermode=count -coverprofile="$pkg_cov" "$pkg"
        if [ -f "$pkg_cov" ]
        then
            go tool cover -html="$pkg_cov" -o "${pkg_name}.html"
        fi
    done
}

case $1 in
    test)
        go_fmt
        go_vet
        go_test
        ;;
    *)
        usage
        ;;
esac
