package goget

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Vendor runs go get into local vendor directory
func Vendor(pkgs []string) error {
	// Get the current working directory for the project
	// TODO: how does this work if you aren't in the project dir?
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Make a temp workspace
	ws, err := ioutil.TempDir("", "goget")
	if err != nil {
		return err
	}

	defer func() {
		// TODO: Remove src symbolic link
		err := os.RemoveAll(ws)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create vendor if it doesn't exist
	v := filepath.Join(wd, "vendor")
	err = os.MkdirAll(v, 0755)
	if err != nil {
		return err
	}

	// Symlink vendor/src into temp workspace
	err = os.Symlink(v, filepath.Join(ws, "src"))
	if err != nil {
		return err
	}

	// Run go get -d command
	cmd := exec.Command("go", "get", "-d")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOPATH=%s", ws))
	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
