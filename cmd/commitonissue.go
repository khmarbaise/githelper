package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/urfave/cli/v2"
)

//CommitOnIssue Uses the information from JIRA as a commit message.
var CommitOnIssue = cli.Command{
	Name:        "commitonissue",
	Aliases:     []string{"coi"},
	Usage:       "commit on issue",
	Description: "Commit the current state of branch with commit message related to JIRA issue.",
	Action:      commitonissue,
}

func commitonissue(ctx *cli.Context) error {
	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)
	check.IfError(err)

	if check.IsMainBranch(currentBranch.Branch) {
		return fmt.Errorf("you are currently on %v where your are not allowed to commit", currentBranch.Branch)
	}

	fmt.Printf("%v\n", currentBranch.Branch)
	fmt.Printf("%v\n", currentBranch.Hash)
	//BRANCH=$(git symbolic-ref --short HEAD)
	//if [ $? -ne 0 ]; then
	//  echo "We are not on any branch. (detached?)"
	//  exit 1;
	//fi
	//# If we are already on master it does not make sense
	//# to continue.
	//if [ "$BRANCH" == "master" ]; then
	//  echo "We are on master."
	//  exit 2;
	//fi
	//CHECK_SESSION=$(jira-cli session --quiet)
	//if [ $? -ne 0 ]; then
	//  echo "You are not logged in on JIRA"
	//	jira-cli login
	//fi
	//
	//# TODO: Tweak jira-cli via templates to make this easier.
	//SUMMARY=$(jira-cli view $BRANCH | grep "^summary: " | cut -d " " -f2-)
	//if [ $? -ne 0 ]; then
	//  echo "Failure while getting information from JIRA"
	//  exit 1;
	//fi
	//# commit the curent state.
	//git commit -a -m"[$BRANCH] - $SUMMARY"
	return nil
}
