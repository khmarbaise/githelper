package cmd

import (
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
