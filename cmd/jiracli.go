package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/jira"
	"github.com/urfave/cli/v2"
)

//JiraCli Just calling jira-cli for testing purposes.
var JiraCli = cli.Command{
	Name:        "jira",
	Aliases:     []string{"jira"},
	Usage:       "current",
	Description: "Testing",
	Action:      jiracli,
}

func jiracli(ctx *cli.Context) error {

	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)

	currentBranch, err := modules.GetCurrentBranch(gitRepo)

	check.IfError(err)
	if check.IsMainBranch(currentBranch.Branch) {
		fmt.Println("you are on main branch.")
	}

	jira.Session()

	summary := jira.IssueSummary(currentBranch.Branch)

	fmt.Printf("Jira summary: '%v'\n", summary)
	return nil
}
