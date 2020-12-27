// Copyright 2020 The GitHelper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Drone Settings is command line tool for githelper.
package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

// Version holds the current tea version
var Version = "development"

// Tags holds the build tags used
var Tags = ""

func main() {
	app := cli.NewApp()
	app.Name = "githelper"
	app.Usage = "Git Helper."
	app.Version = Version + formatBuiltWith(Tags)
	//app.Commands = []*cli.Command{
	//	&cmd.CmdTree,
	//}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}

	firstGitFunction()
}

func firstGitFunction() {
	r, err := git.PlainOpen(".")
	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)

	fmt.Printf("Head Reference: name: %v type: %v hash: %v strings:%v", ref.Name(), ref.Type(), ref.Hash(), ref.Strings())
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits
	var cCount int
	err = cIter.ForEach(func(c *object.Commit) error {
		cCount++
		return nil
	})
	CheckIfError(err)

	fmt.Println(cCount)}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func formatBuiltWith(Tags string) string {
	if len(Tags) == 0 {
		return ""
	}

	return " built with: " + strings.Replace(Tags, " ", ", ", -1)
}
