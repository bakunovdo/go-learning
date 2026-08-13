package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// countCommitsByDay returns map[yyyy-mm-dd] -> commit count for given author emails.
func countCommitsByDay(repos []string, emails []string) (map[string]int, error) {
	allowed := make(map[string]struct{}, len(emails))
	for _, e := range emails {
		allowed[strings.ToLower(strings.TrimSpace(e))] = struct{}{}
	}

	counts := make(map[string]int)

	for _, repoPath := range repos {
		n, err := countRepoCommits(repoPath, allowed)
		if err != nil {
			// пустой репо / битый HEAD — просто пропускаем
			if isIgnorableRepoErr(err) {
				continue
			}
			fmt.Fprintf(os.Stderr, "warn: %s: %v\n", repoPath, err)
			continue
		}
		for day, c := range n {
			counts[day] += c
		}
	}

	return counts, nil
}

func countRepoCommits(repoPath string, allowed map[string]struct{}) (map[string]int, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	cIter, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, fmt.Errorf("log: %w", err)
	}

	counts := make(map[string]int)
	err = cIter.ForEach(func(c *object.Commit) error {
		email := strings.ToLower(c.Author.Email)
		if _, ok := allowed[email]; !ok {
			return nil
		}
		day := c.Author.When.Format("2006-01-02")
		counts[day]++
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("iterate: %w", err)
	}

	return counts, nil
}

func isIgnorableRepoErr(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "reference not found") ||
		strings.Contains(msg, "repository does not exist")
}
