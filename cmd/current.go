package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/urfave/cli/v2"
)

//Current prints out some information about the current branch.
var Current = cli.Command{
	Name:        "current",
	Aliases:     []string{"cr"},
	Usage:       "current",
	Description: "Print out information about the current branch.",
	Action:      current,
}

func current(ctx *cli.Context) error {
	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)

	check.IfError(err)

	fmt.Printf(" name: %v\n", currentBranch.Branch)
	fmt.Printf(" hash: %v\n", currentBranch.Hash)

	return nil
}
