package services

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/semgrep-integrator/models"
)

var (
	semgrep models.Semgrep
)

func Jira(opts *models.Options) {
	client, err := jiraAuth(opts)
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error create issue")
		return
	}

	createIssue(client)

}

func createIssue(client *jira.Client) {
	sc := bufio.NewScanner(os.Stdin)
	var semgrep models.Semgrep
	for sc.Scan() {
		json.Unmarshal([]byte(sc.Text()), &semgrep)
		for _, data := range semgrep.Results {
			i := jira.Issue{
				Fields: &jira.IssueFields{
					Description: models.CreateSummaryJira(&data),
					Type: jira.IssueType{
						Name: "Task",
					},
					Project: jira.Project{
						Key: "ST",
					},
					Summary: data.Extra.Metadata.Issue,
				},
			}
			issue, _, err := client.Issue.Create(&i)
			if err != nil {
				gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error create issue")
				continue
			}
			gologger.Info().Str("Is Success ", fmt.Sprintf("%s", issue.Key)).Msg("Success create issue")
		}

	}

}

func jiraAuth(opts *models.Options) (client *jira.Client, err error) {
	// Create a BasicAuth Transport object
	tp := jira.BasicAuthTransport{
		Username: opts.JiraUser,
		Password: opts.JiraToken,
	}
	// Create a new Jira Client
	client, err = jira.NewClient(tp.Client(), opts.JiraURI)
	if err != nil {
		return client, errors.New("Invalid credentials")
	}
	return client, nil
}
