package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
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
	CheckIfError(err)

	ref, err := gitRepo.Head()
	CheckIfError(err)

	fmt.Printf("    name: %v\n", ref.Name().String())
	fmt.Printf("isBranch: %v\n", ref.Name().IsBranch())
	fmt.Printf("isRemote: %v\n", ref.Name().IsRemote())
	fmt.Printf("    type: %v\n", ref.Type())
	fmt.Printf("    hash: %v\n", ref.Hash())

	return nil
}
