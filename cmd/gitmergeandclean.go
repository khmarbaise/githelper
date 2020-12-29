package cmd

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"github.com/urfave/cli/v2"
	"strings"
)

var (
	//GitMergeAndClean The execution of the git merge and clean.
	GitMergeAndClean = cli.Command{
		Name:        "gitmergeandclean",
		Aliases:     []string{"gmc"},
		Usage:       "git merge and clean.",
		Description: "Merge current branch via fast-forward into master.",
		Action:      mergeAndClean,
	}

	//ErrorPleaseCommitYourChange Is thrown if you have changes locally and not committed.
	ErrorPleaseCommitYourChange = errors.New("please commit your changes or stash them before you switch branches")
)

func mergeAndClean(ctx *cli.Context) error {
	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)

	check.IfError(err)

	if check.IsMainBranch(currentBranch.Branch) {
		return fmt.Errorf("you are currently on %v which you can not merge", currentBranch.Branch)
	}

	fmt.Printf("Branch name: %v\n", currentBranch.Branch)
	fmt.Printf("Branch hash: %v\n", currentBranch.Hash)

	branches, err := gitRepo.Branches()
	check.IfError(err)
	var branchNames []string
	_ = branches.ForEach(func(branch *plumbing.Reference) error {
		fmt.Printf(" -> %v hash:%v type:%v \n", branch.Name(), branch.Hash(), branch.Type())
		branchNames = append(branchNames, strings.TrimPrefix(branch.Name().String(), modules.BranchPrefix))
		return nil
	})

	for _, branch := range branchNames {
		fmt.Printf("Branch: '%v'\n", branch)
	}

	worktree, err := gitRepo.Worktree()
	check.IfError(err)

	status, err := worktree.Status()
	check.IfError(err)
	if !status.IsClean() {
		fmt.Println("Status: **NOT CLEAN**")
		return ErrorPleaseCommitYourChange
	}

	//branchRef := plumbing.NewBranchReferenceName("master")

	remote, err := gitRepo.Remote(currentBranch.Branch)
	if err == nil {
		fmt.Printf("Remote: %v\n", remote.Config())
	} else {
		//TODO: Reconsider: remote does not exist ! => Failure?
		fmt.Printf("Remote branch %v not found  %v\n", currentBranch.Branch, err)
		return err
	}
	fmt.Printf("Checking out %v...", "master")
	//checkoutOptions := git.CheckoutOptions{Branch: branchRef, Create: false, Force: true, Keep: false}
	//err = worktree.Checkout(&checkoutOptions)

	//TODO: We should check for either master/main and use the one we found.
	// modules.SearchForBranch(...)
	execute.ExternalCommand("git", "checkout", "master")

	fmt.Printf("\n")

	return nil
}
