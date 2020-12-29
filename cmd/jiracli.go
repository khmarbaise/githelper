package cmd

import (
	"fmt"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
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

const noSessionExists = 1

func jiracli(ctx *cli.Context) error {

	b, err := execute.ExternalCommandWithRedirect("jira-cli", "session", "--quiet")
	if b.ExitCode == noSessionExists && err != nil {
		_, err := execute.ExternalCommandInteractive("jira-cli", "login")
		check.IfError(err)
	}

	b, err = execute.ExternalCommandWithRedirect("jira-cli", "session", "--quiet")
	fmt.Printf("    Code: %v\n", b.ExitCode)
	fmt.Printf("  Stdout: %v\n", b.Stdout)
	fmt.Printf("  Stderr: %v\n", b.Stderr)

	return nil
}
