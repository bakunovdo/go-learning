package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const defaultRoot = "w"
const defaultReposFile = "save.txt"
const defaultEmails = "bakunov.do@gmail.com"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "scan":
		if err := runScan(args); err != nil {
			fmt.Fprintln(os.Stderr, "scan:", err)
			os.Exit(1)
		}
	case "stats":
		if err := runStats(args); err != nil {
			fmt.Fprintln(os.Stderr, "stats:", err)
			os.Exit(1)
		}
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `usage:
  go run . scan  [-root=w] [-file=save.txt]
  go run . stats [-file=save.txt] [-emails=a@b.c,d@e.f]

`)
}

func runScan(args []string) error {
	fs := flag.NewFlagSet("scan", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", defaultRoot, "folder under home to scan")
	file := fs.String("file", defaultReposFile, "output file for repo paths")
	if err := fs.Parse(args); err != nil {
		return err
	}

	repos, err := scanRepos(*root)
	if err != nil {
		return err
	}

	if err := saveRepos(*file, repos); err != nil {
		return err
	}

	fmt.Printf("saved %d repos to %s\n", len(repos), *file)
	return nil
}

func runStats(args []string) error {
	fs := flag.NewFlagSet("stats", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	file := fs.String("file", defaultReposFile, "file with repo paths")
	emailsFlag := fs.String("emails", defaultEmails, "comma-separated author emails")
	if err := fs.Parse(args); err != nil {
		return err
	}

	emails := splitEmails(*emailsFlag)
	if len(emails) == 0 {
		return fmt.Errorf("no emails provided (use -emails=a@b.c)")
	}

	repos, err := loadRepos(*file)
	if err != nil {
		return fmt.Errorf("read %s (run scan first): %w", *file, err)
	}

	counts, err := countCommitsByDay(repos, emails)
	if err != nil {
		return err
	}

	printBoard(counts)
	return nil
}

func splitEmails(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
