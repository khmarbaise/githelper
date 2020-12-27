// Copyright 2020 The GitHelper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Drone Settings is command line tool for githelper.
package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
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

// branchPrefix base dir of the branch information file store on git
const branchPrefix = "refs/heads/"

func firstGitFunction() {
	gitRepo, err := git.PlainOpen(".")
	CheckIfError(err)

	ref, err := gitRepo.Head()
	CheckIfError(err)

	fmt.Printf("Head Reference: name: %v type: %v hash: %v strings:%v\n", ref.Name(), ref.Type(), ref.Hash(), ref.Strings())

	if !strings.HasPrefix(ref.Name().String(), branchPrefix) {
		fmt.Errorf("invalid HEAD branch: %v", ref.String())
	}

	branch := ref.Name().String()[len(branchPrefix):]

	//if branch != "main" && branch != "master" {
	//	fmt.Errorf("We are main/master.", branch)
	//}

	fmt.Printf("Branch name: %v\n", branch)

	branches, err := gitRepo.Branches()
	CheckIfError(err)
	var branchNames []string
	_ = branches.ForEach(func(branch *plumbing.Reference) error {
		fmt.Printf(" -> %v hash:%v type:%v \n", branch.Name(), branch.Hash(), branch.Type())
		branchNames = append(branchNames, strings.TrimPrefix(branch.Name().String(), branchPrefix))
		return nil
	})

	for _, branch := range branchNames {
		fmt.Printf("Branch: '%v'\n", branch)
	}

	worktree, err := gitRepo.Worktree()
	CheckIfError(err)

	status, err := worktree.Status()
	CheckIfError(err)
	if !status.IsClean() {
		fmt.Println("Status: **NOT CLEAN**")
		fmt.Println("Please commit your changes or stash them before you switch branches.")
		// stop working because we can't change the branch.
		return
	}

	//branchRef := plumbing.NewBranchReferenceName("master")

	remote, err := gitRepo.Remote(ref.Name().String())
	if err == nil {
		fmt.Printf("Remote: %v\n", remote.Config())
	} else {
		//TODO: Reconsider: remote does not exist ! => Failure?
		fmt.Printf("Remote branch %v not found  %v\n", ref.Name(), err)
		return
	}
	fmt.Printf("Checking out %v...", "master")
	//TODO: We should check for either master/main and use the one we found.
	//checkoutOptions := git.CheckoutOptions{Branch: branchRef, Create: false, Force: true, Keep: false}
	//err = worktree.Checkout(&checkoutOptions)
	runCmd("git", "checkout", "master")

	fmt.Printf("\n")
}

func runCmd(cmd ...string) {
	log.Printf("Executing : %s ...\n", cmd)
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		log.Panicln(err)
	}
	if err := c.Wait(); err != nil {
		log.Panicln(err)
	}
}

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
