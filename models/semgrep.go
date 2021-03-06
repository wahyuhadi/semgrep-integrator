package models

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Semgrep struct {
	Errors  []interface{} `json:"errors"`
	Results []struct {
		CheckID string `json:"check_id"`
		End     struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"end"`
		Extra struct {
			IsIgnored bool   `json:"is_ignored"`
			Lines     string `json:"lines"`
			Message   string `json:"message"`
			Metadata  struct {
				Issue         string   `json:"issue"`
				Severity      string   `json:"severity"`
				Category      string   `json:"category"`
				Cwe           string   `json:"cwe"`
				Impact        []string `json:"impact"`
				Owasp         string   `json:"owasp"`
				References    []string `json:"references"`
				Remediation   []string `json:"remediation"`
				SourceRuleURL string   `json:"source-rule-url"`
				Technology    []string `json:"technology"`
			} `json:"metadata"`
			Metavars struct {
			} `json:"metavars"`
			Severity string `json:"severity"`
		} `json:"extra"`
		Path  string `json:"path"`
		Start struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"start"`
	} `json:"results"`
}

type Results struct {
	CheckID string `json:"check_id"`
	End     struct {
		Col    int `json:"col"`
		Line   int `json:"line"`
		Offset int `json:"offset"`
	} `json:"end"`
	Extra struct {
		IsIgnored bool   `json:"is_ignored"`
		Lines     string `json:"lines"`
		Message   string `json:"message"`
		Metadata  struct {
			Issue         string   `json:"issue"`
			Severity      string   `json:"severity"`
			Category      string   `json:"category"`
			Cwe           string   `json:"cwe"`
			Impact        []string `json:"impact"`
			Owasp         string   `json:"owasp"`
			References    []string `json:"references"`
			Remediation   []string `json:"remediation"`
			SourceRuleURL string   `json:"source-rule-url"`
			Technology    []string `json:"technology"`
		} `json:"metadata"`
		Metavars struct {
		} `json:"metavars"`
		Severity string `json:"severity"`
	} `json:"extra"`
	Path  string `json:"path"`
	Start struct {
		Col    int `json:"col"`
		Line   int `json:"line"`
		Offset int `json:"offset"`
	} `json:"start"`
}

func CreateTemplateGithubIssue(parse interface{}) (body string) {
	var data Results
	mapstructure.Decode(parse, &data)
	issue := fmt.Sprintf("**Issue : %s**\n ", data.Extra.Metadata.Issue)
	severity := fmt.Sprintf("**Severity : %s**\n ", data.Extra.Metadata.Severity)
	filename := fmt.Sprintf("**Filename:** %s:%v:%v\n ", data.Path, data.Start.Col, data.Start.Line)
	impact := fmt.Sprintf("**Impact :** %s", data.Extra.Metadata.Impact)
	body = issue + severity + filename + impact
	return body
}

func CreateSummaryJira(data interface{}) (summ string) {
	var semgrep Results
	mapstructure.Decode(data, &semgrep)
	Desc := fmt.Sprintf("Desc : %s\n", semgrep.Extra.Message)
	Impact := fmt.Sprintf("Impact : %s\n", semgrep.Extra.Metadata.Impact)
	Path := fmt.Sprintf("Path : %v:%v:%v\n", semgrep.Path, semgrep.Start.Line, semgrep.Start.Col)
	Severity := fmt.Sprintf("Severity : %s\n", semgrep.Extra.Metadata.Severity)

	summ = Desc + Impact + Path + Severity
	fmt.Println(summ)
	return summ
}
