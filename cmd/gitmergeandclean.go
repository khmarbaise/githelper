package cmd

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
	"github.com/khmarbaise/githelper/modules/execute"
	"github.com/khmarbaise/githelper/modules/githelper"
	"github.com/khmarbaise/githelper/modules/jira"
	"github.com/urfave/cli/v2"
)

var (
	//GitMergeAndClean The execution of the git merge and clean.
	GitMergeAndClean = cli.Command{
		Name:        "gitmergeandclean",
		Aliases:     []string{"gmc"},
		Usage:       "git merge and clean.",
		Description: "Merge current branch via fast-forward into master and delete the remote branch and afterwards the local branch as well..",
		Action:      mergeAndClean,
	}

	//ErrorPleaseCommitYourChange Is thrown if you have changes locally and not committed.
	ErrorPleaseCommitYourChange = errors.New("please commit your changes or stash them before you switch branches")
)

func mergeAndClean(ctx *cli.Context) error {
	fmt.Printf("%v...", "Opening repository")
	gitRepo, err := git.PlainOpen(".")
	check.IfError(err)
	fmt.Println("done.")

	currentBranch, err := modules.GetCurrentBranch(gitRepo)

	check.IfError(err)

	if check.IsMainBranch(currentBranch.Branch) {
		return fmt.Errorf("you are currently on %v which you can not merge", currentBranch.Branch)
	}

	mainBranch, err := modules.SearchForMainBranch(gitRepo)
	check.IfError(err)

	worktree, err := gitRepo.Worktree()
	check.IfError(err)

	fmt.Printf("%v...", "Check for changes...")
	status, err := worktree.Status()
	check.IfError(err)
	if !status.IsClean() {
		fmt.Println("Status: **NOT CLEAN**")
		return ErrorPleaseCommitYourChange
	}
	fmt.Println("done.")

	fmt.Printf("Checking out '%v'...", mainBranch)
	b, err := execute.ExternalCommandWithRedirect("git", "checkout", mainBranch)
	check.IfErrorWithOutput(err, b.Stdout, b.Stderr)
	fmt.Println("done.")

	fmt.Printf("Merging '%v' into '%v' via fast forward only...", currentBranch.Branch, mainBranch)
	b, err = execute.ExternalCommandWithRedirect("git", "merge", "--ff-only", currentBranch.Branch)
	check.IfErrorWithOutput(err, b.Stdout, b.Stderr)
	fmt.Println("done.")

	fmt.Printf("Push '%v' to remote...", mainBranch)
	b, err = execute.ExternalCommandWithRedirect("git", "push", "origin", mainBranch)
	check.IfErrorWithOutput(err, b.Stdout, b.Stderr)
	fmt.Println("done.")

	fmt.Printf("Delete remote branch '%v'...", currentBranch.Branch)
	b, err = execute.ExternalCommandWithRedirect("git", "push", "origin", "--delete", currentBranch.Branch)
	check.IfErrorWithOutput(err, b.Stdout, b.Stderr)
	fmt.Println("done.")

	fmt.Printf("Delete local branch '%v'...", currentBranch.Branch)
	// We assume that the merge has been done successfully otherwise this will fail.
	b, err = execute.ExternalCommandWithRedirect("git", "branch", "-d", currentBranch.Branch)
	check.IfErrorWithOutput(err, b.Stdout, b.Stderr)
	fmt.Println("done.")

	uri := githelper.GetGitRemoteURI(gitRepo)

	//TODO: Reconsider to handle this via a configuration.
	switch uri.Host {
	case "github.com":
		fmt.Printf("--> Github.com\n")
		// We don't need to do something because an issue will be closed automatically
		// via "Fixed #" in commit message.
		break
	case "gitea.com":
		fmt.Printf("--> Gitea.com\n")
		// We don't need to do something because an issue will be closed automatically
		// via "Fixed #" in commit message.
		break
	case "gitbox.apache.org":
		fmt.Printf("--> Apache Gitbox\n")
		jira.Session()
		link := fmt.Sprintf("%v://%v%v?p=%v;h=%v", uri.Schema, uri.Host, uri.Base, uri.Project, currentBranch.Hash)
		commitMessage := fmt.Sprintf("Done in [%v|%v]", currentBranch.Hash, link)
		fmt.Printf("Closing issue %v...", currentBranch.Branch)
		execute.ExternalCommandWithRedirect("jira-cli", "close", "-m", commitMessage, "--resolution=Done", currentBranch.Branch)
		fmt.Println("Done.")
		break
	default:
		return fmt.Errorf("unknown host %v used", uri.Host)
	}

	return nil
}
