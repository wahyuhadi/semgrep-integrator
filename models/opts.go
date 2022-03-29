package models

type Options struct {
	GithubIssue bool
	GithubToken string
	Owner       string
	RepoName    string

	Elasic      bool
	ElasticUrl  string
	ElasticUser string
	ElasticPass string
	ElasicIndex string

	Jira      bool
	JiraUser  string
	JiraToken string
	JiraURI   string
}
