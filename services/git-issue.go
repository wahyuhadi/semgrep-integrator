package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/semgrep-integrator/models"
	"golang.org/x/oauth2"
)

func CreateGithubIssue(opts *models.Options) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: opts.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var semgrep models.Semgrep
	err := json.NewDecoder(os.Stdin).Decode(&semgrep)
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error json decoder")
	}
	if len(semgrep.Results) > 0 {
		for _, data := range semgrep.Results {
			body := models.CreateTempalte(data)
			input := &github.IssueRequest{
				Title:  github.String(data.Extra.Metadata.Issue),
				Body:   github.String(body),
				Labels: &[]string{"Bug", "Security"},
			}
			issue, _, err := client.Issues.Create(ctx, opts.Owner, opts.RepoName, input)
			if err != nil {
				gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error when create issue")
			}
			gologger.Info().Str("Issue", fmt.Sprintf("%s", data.Extra.Metadata.Issue)).Str("Issue Number", fmt.Sprintf("%v", *issue.Number)).Msg("Success Create issue on github repo")
		}
	}
}
