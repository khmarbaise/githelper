package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"github.com/urfave/cli/v2"
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

	if check.IsMainBranch(currentBranch.Branch) {
		return fmt.Errorf("you are currently on %v which is not allowed to be force pushed", currentBranch.Branch)
	}

	//FIXME: If we do push the first time setup tracking branch as well.
	r, err := execute.ExternalCommandWithRedirect("git", "push", "origin", "--force-with-lease", currentBranch.Branch)
	check.IfErrorWithOutput(err, r.Stdout, r.Stderr)

	return nil
}
