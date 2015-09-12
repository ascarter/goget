package goget

/*

#!/bin/sh

workingd=`mktemp -d -t goget` || exit 1

# Make sure vendor exists
# TODO

# Use workingd as the workspace and set src to be vendor
ln -s `pwd`/vendor ${workingd}/src

GOPATH=${workingd} go get -d -u $*
#ls -l ${workingd}

# Remove link
rm ${workingd}/src

*/

func Get(pkgs []string) {
	// Make a temp workspace
	// Create vendor if it doesn't exist
	// Symlink vendor/src into temp workspace
	// Run go get -d command
	// Cleanup
}
