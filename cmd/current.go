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
	if check.IsMainBranch(currentBranch.Branch) {
		fmt.Println("you are on main branch.")
	}

	remotes, err := gitRepo.Remotes()
	check.IfError(err)

	for _, remote := range remotes {
		fmt.Printf("Remote: '%v'\n", remote.String())
		fmt.Printf("config: '%v'\n", remote.Config().Name)
		fmt.Printf("fetch: '%v'\n", remote.Config().Fetch)
		fmt.Println("--- Fetch --- ")
		for _, f := range remote.Config().Fetch {
			fmt.Printf(" ->        string: '%v'\n", f.String())
			fmt.Printf(" ->      IsDelete: '%v'\n", f.IsDelete())
			fmt.Printf(" ->       Reverse: '%v'\n", f.Reverse())
			fmt.Printf(" -> IsForceUpdate: '%v'\n", f.IsForceUpdate())
			fmt.Printf(" ->           src: '%v'\n", f.Src())
		}
		fmt.Println("--- URLs --- ")
		for _, url := range remote.Config().URLs {
			fmt.Printf(" ->     url: '%v'\n", url)
		}
	}

	fmt.Printf(" name: %v\n", currentBranch.Branch)
	fmt.Printf(" hash: %v\n", currentBranch.Hash)

	return nil
}
