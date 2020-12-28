package modules

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"strings"
)

//SearchForBranch Will search for the Branch names "master" and "main".
func SearchForBranch(repository *git.Repository) error {
	return nil
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
