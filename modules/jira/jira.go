package jira

import (
	"fmt"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"strings"
)

const noSessionExists = 1

//JiraSession Will check if a session for jira-cli exists or not. If not it will call jira-cli interactive
// to create an appropriate session.
func JiraSession() {
	b, err := execute.ExternalCommandWithRedirect("jira-cli", "session", "--quiet")
	if b.ExitCode == noSessionExists && err != nil {
		fmt.Println("There is not existing session for jira.")
		_, err := execute.ExternalCommandInteractive("jira-cli", "login")
		check.IfError(err)
	}
}

//JiraIssueSummary Call "jira-cli view $BRANCH" and extract the line "^summary: (.*)"
func JiraIssueSummary(branch string) string {
	b, err := execute.ExternalCommandWithRedirect("jira-cli", "view", branch)
	check.IfError(err)
	split := strings.Split(strings.Replace(b.Stdout, "\r\n", "\n", -1), "\n")

	return modules.ExtractSummary(split)
}
