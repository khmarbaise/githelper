package cmd

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"github.com/urfave/cli/v2"
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

	mainBranch, err := modules.SearchForMainBranch(gitRepo)
	check.IfError(err)

	worktree, err := gitRepo.Worktree()
	check.IfError(err)

	status, err := worktree.Status()
	check.IfError(err)
	if !status.IsClean() {
		fmt.Println("Status: **NOT CLEAN**")
		return ErrorPleaseCommitYourChange
	}

	//branchRef := plumbing.NewBranchReferenceName("master")

	fmt.Printf("Checking out %v\n", mainBranch)
	execute.ExternalCommand("git", "checkout", mainBranch)

	fmt.Printf("Merging %v into %v via fast forward only\n", currentBranch.Branch, mainBranch)
	execute.ExternalCommand("git", "merge", "--ff-only", currentBranch.Branch)

	fmt.Printf("Push %v to remote\n", mainBranch)
	execute.ExternalCommand("git", "push", "origin", mainBranch)

	fmt.Printf("Delete remote %v\n", currentBranch.Branch)
	execute.ExternalCommand("git", "push", "origin", "--delete", currentBranch.Branch)

	fmt.Printf("Delete local %v \n", currentBranch.Branch)
	// We assume that the merge has been done successfully otherwise this will fail.
	execute.ExternalCommand("git", "branch", "-d", currentBranch.Branch)

	return nil

	//# Get the latest commit HASH
	//#
	//COMMITHASH=$(git rev-parse HEAD)
	//#
	//# Get the GIT URL from pom file:
	//# TODO: Can we do some sanity checks? Yes: scm:git:..  if not FAIL!
	//echo -n "Get the git url from pom file..."
	//GITURL=$(mvn org.apache.maven.plugins:maven-help-plugin:3.2.0:evaluate -Dexpression=project.scm.connection -q -DforceStdout | cut -d":" -f3-)
	//echo " '$GITURL' done."
	//GITPROJECT=$(basename $GITURL)
	//GITBASE=$(dirname $GITURL)
	//#
	//#
	//# Check if we are github project => GitHub issue tracker
	//# Check if we are gitbox project => JIRA issue tracker
	//#    We extracting 1. github.com
	//#                  2. gitbox.apache.org
	//GITHOST=$(echo $GITURL | cut -d ":" -f2- | cut -d "/" -f3 )
	//if [ "$GITHOST" == "github.com" ]; then
	//	echo "GitHub Issue Tracker"
	//	exit 0;
	//else
	//	echo "JIRA Issue Tracker (Apache Project)"
	//fi
	//#
	//CHECK_SESSION=$(jira-cli session --quiet)
	//if [ $? -ne 0 ]; then
	//  echo "You are not logged in on JIRA"
	//	jira-cli login
	//fi
	//#
	//echo "Closing JIRA issue $BRANCH"
	//jira-cli close -m"Done in [$COMMITHASH|$GITBASE?p=$GITPROJECT;a=commitdiff;h=$COMMITHASH]" --resolution=Done $BRANCH
	//## Error handling?
	//echo "Closing finished."
	//#

}
