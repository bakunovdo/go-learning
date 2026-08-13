package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func scanRepos(rootUnderHome string) ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("home dir: %w", err)
	}

	root := filepath.Join(home, rootUnderHome)
	var repos []string

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "warn: skip %s: %v\n", path, err)
			return nil
		}
		if !d.IsDir() {
			return nil
		}

		name := d.Name()
		if name == "node_modules" || name == ".git" || name == "vendor" {
			return filepath.SkipDir
		}

		gitPath := filepath.Join(path, ".git")
		if _, err := os.Stat(gitPath); err != nil {
			return nil
		}

		repos = append(repos, path)
		// не обходим содержимое уже найденного репозитория
		return filepath.SkipDir
	})
	if err != nil {
		return nil, fmt.Errorf("walk %s: %w", root, err)
	}

	return repos, nil
}

func saveRepos(file string, repos []string) error {
	data := strings.Join(repos, "\n")
	if len(repos) > 0 {
		data += "\n"
	}
	if err := os.WriteFile(file, []byte(data), 0644); err != nil {
		return fmt.Errorf("write %s: %w", file, err)
	}
	return nil
}

func loadRepos(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	repos := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			repos = append(repos, line)
		}
	}
	return repos, nil
}
