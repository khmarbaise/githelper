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

	fmt.Printf("Branch name: #{currentBranch.Branch}\n")
	fmt.Printf("Branch hash: %v\n", currentBranch.Hash)

	branch, err := modules.SearchForMainBranch(gitRepo)
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

	execute.ExternalCommand("git", "checkout", branch)

	fmt.Printf("\n")

	return nil

	//# If we are on another branch goto master
	//git co master
	//# Only allow fast-forward merges..
	//git merge --ff-only $BRANCH
	//if [ $? -ne 0 ]; then
	//  echo "git merge can't be fast forwarded"
	//  exit 1;
	//fi
	//git push origin master
	//if [ $? -ne 0 ]; then
	//  echo "git push to master has failed. rejected?"
	//  exit 1;
	//fi
	//# Improvement
	//#  Try to identify if a branch is remotely being tracked? or exists remotely?
	//#
	//# delete remote branch
	//git push origin --delete $BRANCH
	//if [ $? -ne 0 ]; then
	//  echo "failed to delete $BRANCH ?"
	//  exit 1;
	//fi
	//# delete local branch
	//git branch -d $BRANCH
	//#
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
