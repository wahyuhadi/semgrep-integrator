package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/semgrep-integrator/models"
	"github.com/wahyuhadi/semgrep-integrator/services"
)

var (
	issue = flag.Bool("gi", false, "Create issue to github")
	owner = flag.String("o", "", "Owner or Org repo")
	repo  = flag.String("r", "", "Repository name")
)

func parseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		GithubIssue: *issue,
		Owner:       *owner,
		RepoName:    *repo,
		GithubToken: os.Getenv("semgrep_integ"),
	}
}

func main() {
	opts := parseOptions()
	if opts.GithubToken == "" {
		gologger.Info().Str("Error", fmt.Sprintf("%v", "no gittoken found")).Msg("Add semgrep_integ in .zshrc")
		return
	}
	// report to github issue
	if opts.GithubIssue {
		services.CreateGithubIssue(opts)
	}

	// todo something magic
	// Jira
	// Elastic
	// Slack
}
