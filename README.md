# Semgrep-integrator

### Scan your code with custom rules and integrate it with anything, anytime, anywhere

----
- Example Gorm SQL-Inection rule with .yml
```yml
# yamllint disable-line rule:document-start
rules:
# yamllint disable-line rule:indentation
- id: SQL-Injection
  pattern: $W.Where("id = ", ...)
  message: Posible SQL Injection
  languages: [go]
  severity: WARNING
  metadata:
      issue: "SQL Injection"
      severity: "HIGH"
      owasp: "SQL Injection"
      category: security
      impact: "Potential Data Breach"
      # yamllint disable-line rule:new-line-at-end-of-file
      remediation: "Validation user input"
```

# Will be integrate with:
- Github
- Jira
- Elastic
- Slack 
- Etc..

![sc](/img.png)