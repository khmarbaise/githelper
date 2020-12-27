package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/khmarbaise/githelper/modules"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

//GitMergeAndClean The execution of the git merge and clean.
var GitMergeAndClean = cli.Command{
	Name:        "gitmergeandclean",
	Aliases:     []string{"gmc"},
	Usage:       "git merge and clean.",
	Description: "Merge current branch via fast-forward into master.",
	Action:      mergeAndClean,
}

// branchPrefix base dir of the branch information file store on git
const branchPrefix = "refs/heads/"

func mergeAndClean(ctx *cli.Context) error {
	gitRepo, err := git.PlainOpen(".")
	CheckIfError(err)

	ref, err := gitRepo.Head()
	CheckIfError(err)

	fmt.Printf("Head Reference: name: %v type: %v hash: %v strings:%v\n", ref.Name(), ref.Type(), ref.Hash(), ref.Strings())

	if !strings.HasPrefix(ref.Name().String(), branchPrefix) {
		fmt.Errorf("invalid HEAD branch: %v", ref.String())
	}

	branch := ref.Name().String()[len(branchPrefix):]
	branchHash := ref.Hash()

	//FIXME: Check for main/master
	//if branch != "main" && branch != "master" {
	//	fmt.Errorf("We are main/master.", branch)
	//}

	fmt.Printf("Branch name: %v\n", branch)
	fmt.Printf("Branch hash: %v\n", branchHash)

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
		//FIXME: return a real error!!!
		return nil
	}

	// Need more.
	//branchRef := plumbing.NewBranchReferenceName("master")

	remote, err := gitRepo.Remote(ref.Name().String())
	if err == nil {
		fmt.Printf("Remote: %v\n", remote.Config())
	} else {
		//TODO: Reconsider: remote does not exist ! => Failure?
		fmt.Printf("Remote branch %v not found  %v\n", ref.Name(), err)
		return err
	}
	fmt.Printf("Checking out %v...", "master")
	//checkoutOptions := git.CheckoutOptions{Branch: branchRef, Create: false, Force: true, Keep: false}
	//err = worktree.Checkout(&checkoutOptions)

	//TODO: We should check for either master/main and use the one we found.
	// modules.SearchForBranch(...)
	modules.RunExternalGit("git", "checkout", "master")

	fmt.Printf("\n")

	return nil
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
