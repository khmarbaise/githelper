// Copyright 2020 The GitHelper Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// git-jira-helper is command line tool to support the handling of branches and merges with feature branches.
package main

import (
	"fmt"
	"github.com/khmarbaise/githelper/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

// Version holds the current gjh version
var Version = "development"

// Tags holds the build tags used
var Tags = ""

func main() {
	app := cli.NewApp()
	app.Name = "gjh"
	app.Usage = "Command line tool to support git/jira/issue tracking systems."
	app.Description = "Git Jira Helper. Simpler handling of branches / jira support / issue tracking."
	app.Version = Version + formatBuiltWith(Tags)
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&cmd.GitMergeAndClean,
		&cmd.GitPushWithLease,
		&cmd.Current,
		&cmd.CommitOnIssue,
		&cmd.JiraCli,
		&cmd.StartIssue,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}

}

func formatBuiltWith(Tags string) string {
	if len(Tags) == 0 {
		return ""
	}

	return " built with: " + strings.Replace(Tags, " ", ", ", -1)
}
