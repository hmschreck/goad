package main

import (
	"github.com/akamensky/argparse"
	"github.com/go-git/go-git/v5"
)

var UseGitCommit *bool = parser.Flag("g", "git", &argparse.Options{
	Required: false,
	Default:  false,
	Help:     "use Git data to tag the output",
})

var GitCommit *string = parser.String("", "git-commit", &argparse.Options{
	Required: false,
	Help:     "Git commit to use to tag the output (if none specified, you should use git-path",
	Default:  "",
})

var GitPath *string = parser.String("", "git-path", &argparse.Options{
	Required: false,
	Help:     "use a path to a git repository to get the commit from",
	Default:  ".",
})

// use the flags to return a usable git commit string for performance tracking
func ParseGit() (commit string, err error) {
	if *UseGitCommit {
		if *GitCommit != "" {
			return *GitCommit, nil
		} else {
			repository, err := git.PlainOpen(*GitPath)
			if err != nil {
				return commit, err
			}
			head, err := repository.Head()
			if err != nil {
				return commit, err
			}
			commit = head.Hash().String()
		}
		return
	} else {
		return
	}
}
