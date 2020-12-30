package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/jira"
	"github.com/urfave/cli/v2"
)

//StartIssue Just setting an existing issue into state "in progress"
var StartIssue = cli.Command{
	Name:        "startissue",
	Aliases:     []string{"si"},
	Usage:       "starting the given issue (via branch)",
	Description: "testin",
	Action:      startIssue,
}

func startIssue(ctx *cli.Context) error {

	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)

	check.IfError(err)
	if check.IsMainBranch(currentBranch.Branch) {
		fmt.Println("you are on main branch.")
	}

	jira.Session()

	jira.StartIssue(currentBranch.Branch)

	return nil
}
