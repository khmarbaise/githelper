package githelper

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/khmarbaise/githelper/modules"
	"github.com/khmarbaise/githelper/modules/check"
)

//GetGitRemoteURI Get the remote URI of the current git repository.
func GetGitRemoteURI(gitRepo *git.Repository) (urlParts modules.URLParts) {

	remotes, err := gitRepo.Remotes()
	check.IfError(err)

	urlParts = modules.URLParts{User: "", Schema: "", Host: "", Path: ""}
	for _, remote := range remotes {
		for _, singleURL := range remote.Config().URLs {
			fmt.Printf(" ->     singleURL: '%v'\n", singleURL)

			parse, err := modules.ParseGitURI(singleURL)
			check.IfError(err)
			return parse
		}
	}
	return
}
