package modules

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/khmarbaise/githelper/modules/check"
	"strings"
)

//SearchForMainBranch Will search for the Branch name either "master" or "main".
func SearchForMainBranch(gitRepo *git.Repository) (branch string, err error) {
	branches, err := gitRepo.Branches()
	check.IfError(err)

	var branchNames []string
	_ = branches.ForEach(func(branch *plumbing.Reference) error {
		fmt.Printf(" -> %v hash:%v type:%v \n", branch.Name(), branch.Hash(), branch.Type())
		branchName := strings.TrimPrefix(branch.Name().String(), BranchPrefix)
		if check.IsMainBranch(branchName) {
			branchNames = append(branchNames, branchName)
		}
		return nil
	})

	if len(branchNames) > 1 {
		return "", fmt.Errorf("more than one branch %v found", branchNames)
	}
	return branchNames[0], nil
}

//CurrentBranch Defines information about current branch the name and the hash.
type CurrentBranch struct {
	Branch string
	Hash   plumbing.Hash
}

//GetCurrentBranch Get the current Branch like /refs/head/main and returns "main" incl. Hash code.
// This is the equivalent of "git symbolic-ref --short HEAD".
func GetCurrentBranch(gitRepo *git.Repository) (currentBranch CurrentBranch, err error) {
	currentBranch = CurrentBranch{}
	err = nil

	ref, err := gitRepo.Head()
	if err != nil {
		return
	}

	if !strings.HasPrefix(ref.Name().String(), BranchPrefix) {
		err = fmt.Errorf("invalid HEAD Branch: %v", ref.String())
		return
	}

	return CurrentBranch{Branch: ref.Name().String()[len(BranchPrefix):], Hash: ref.Hash()}, nil
}
