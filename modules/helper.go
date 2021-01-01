package modules

import (
	"fmt"
	"github.com/khmarbaise/githelper/modules/check"
	"net/url"
	"path/filepath"
	"strings"
)

const summaryPrefix = "summary: "

//ExtractSummary Extract a line "^summary: (.*)" from the jira output lines.
func ExtractSummary(lines []string) (result string) {
	result = ""
	for _, line := range lines {
		if strings.HasPrefix(line, summaryPrefix) {
			result = line[len(summaryPrefix):]
		}
	}
	return
}

//URLParts The parts of an git remote url.
type URLParts struct {
	User    string
	Schema  string
	Host    string
	Path    string
	Base    string
	Project string
}

func (u URLParts) String() string {
	return fmt.Sprintf("User: '%v' Schema: '%v' Host: '%v' Path: '%v' Base: '%v' Project: '%v'", u.User, u.Schema, u.Host, u.Path, u.Base, u.Project)
}

//ParseGitURI Will parse a URL which is given by `git remote -v`
//TODO: Reconsider if this is really necessary at all?
func ParseGitURI(uri string) (URLParts, error) {
	parts := URLParts{}
	if strings.HasPrefix(uri, "git@") {
		parts.User = uri[0:3] // extract git only.
		uri = uri[4:]
		index := strings.Index(uri, ":")
		parts.Host = uri[0:index]
		parts.Path = uri[index+1:]
		parts.Schema = "ssh"
		parts.Project = filepath.Base(parts.Path)
		parts.Base = filepath.Dir(parts.Path)
	} else {
		//TODO: Need to reconsider if this is really necessary
		parse, err := url.Parse(uri)
		check.IfError(err)
		parts.User = parse.User.String()
		parts.Schema = parse.Scheme
		parts.Host = parse.Host
		parts.Path = parse.Path
		parts.Project = filepath.Base(parts.Path)
		parts.Base = filepath.Dir(parts.Path)
	}
	return parts, nil
}
