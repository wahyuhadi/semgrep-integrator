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

	epush  = flag.Bool("ep", false, "Elastic push")
	eurl   = flag.String("eurl", "http://127.0.0.1:9200", "Elastic URL")
	euser  = flag.String("euser", "", "elastic users")
	epass  = flag.String("epass", "", "elastic pass")
	eindex = flag.String("eindex", "semgrep-sast", "elastic index")
	jira   = flag.Bool("jira", false, "auto create issue on jira")
)

func parseOptions() (opts *models.Options) {
	flag.Parse()
	return &models.Options{
		GithubIssue: *issue,
		Owner:       *owner,
		RepoName:    *repo,
		GithubToken: os.Getenv("semgrep_integ"),
		Elasic:      *epush,
		ElasticUrl:  *eurl,
		ElasticUser: *euser,
		ElasticPass: *epass,
		ElasicIndex: *eindex,

		// Jira part
		Jira:      *jira,
		JiraUser:  os.Getenv("jira_user"),
		JiraToken: os.Getenv("jira_token"),
		JiraURI:   os.Getenv("jira_url"),
	}
}

func main() {
	opts := parseOptions()

	// report to github issue
	if opts.GithubIssue {
		if opts.GithubToken == "" {
			gologger.Info().Str("Error", fmt.Sprintf("%v", "no gittoken found")).Msg("Add semgrep_integ in .zshrc")
			return
		}
		services.CreateGithubIssue(opts)
	}

	// push data to Elasticsearch
	if opts.Elasic {
		services.Elastic(opts)
	}

	// todo something magic
	// Jira
	// auto create issue in Jira
	if opts.Jira {
		services.Jira(opts)
	}
	// Elastic
	// Slack
}
