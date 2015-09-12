package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// Flags
	downloadOnly = flag.Bool("d", false, "stop after downloading the packages (do not install)")
	forceUpdate  = flag.Bool("f", false, "force update")
	includeTest  = flag.Bool("t", false, "download test packages")
	update       = flag.Bool("u", false, "update named packages and dependencies")
	fix          = flag.Bool("fix", false, "run fix tool before resolving dependencies")
	insecure     = flag.Bool("insecure", false, "permit insecure schemes like HTTP for fetching and resolving")
	// TODO: add go build flags...
	vendor = flag.Bool("v", false, "vendor package")
)

func init() {
	// flag setup

}

func usage() {
	desc := `usage: goget [flags] [packages]

goget downloads and installs packages named by the import paths along with dependencies.

It is a drop-in replacment for 'go get' and extends it to support vendoring in addition
to the workspace. Using '-v' will fetch packages and dependencies to the project 'vendor/'
directory and write the vendor manifest.

Flags:
`
	// get [-d] [-f] [-fix] [-insecure] [-t] [-u] [build flags] [packages]
	fmt.Fprintf(os.Stderr, desc)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	fmt.Println("This is goget")
}
