package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"github.com/urfave/cli/v2"
	"strings"
)

//GitPushWithLease git push with lease.
var GitPushWithLease = cli.Command{
	Name:        "gitpushwithlease",
	Aliases:     []string{"pwl"},
	Usage:       "git push with lease.",
	Description: "Git push with lease (git push --force-with-lease)",
	Action:      pushWithLease,
}

func pushWithLease(ctx *cli.Context) error {
	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)
	check.IfError(err)

	if isMainBranch(currentBranch.Branch) {
		return fmt.Errorf("you are currently on %v which is not alloed to force pushed", currentBranch.Branch)
	}

	execute.ExternalCommand("git", "push", "origin", "--force-with-lease", currentBranch.Branch)
	//TODO:
	//  BRANCH=$(git symbolic-ref --short HEAD)
	//if [ $? -ne 0 ]; then
	//  echo "We are not on any branch. (detached?)"
	//  exit 1;
	//fi
	//if [ "$BRANCH" == "master" ]; then
	//  echo "We are on master."
	//  exit 2;
	//fi
	//git push origin --force-with-lease $BRANCH
	//FIXME: If we do push the first time setup tracking branch as well.
	return nil
}

func isMainBranch(branch string) bool {
	branchWithoutSpaces := strings.TrimSpace(branch)
	if branchWithoutSpaces == "main" || branchWithoutSpaces == "master" {
		return true
	}
	return false
}
