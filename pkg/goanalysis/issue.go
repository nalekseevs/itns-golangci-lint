package goanalysis

import (
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/result"
)

type Issue struct {
	result.Issue
	Pass *analysis.Pass
}

func NewIssue(issue *result.Issue, pass *analysis.Pass) Issue {
	return Issue{
		Issue: *issue,
		Pass:  pass,
	}
}

type EncodingIssue struct {
	FromLinter           string
	Text                 string
	Severity             string
	Pos                  token.Position
	LineRange            *result.Range
	Replacement          *result.Replacement
	ExpectNoLint         bool
	ExpectedNoLintLinter string
}
