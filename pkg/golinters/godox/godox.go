package godox

import (
	"go/token"
	"strings"
	"sync"

	"github.com/matoous/godox"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
	"github.com/nalekseevs/itns-golangci-lint/pkg/lint/linter"
	"github.com/nalekseevs/itns-golangci-lint/pkg/result"
)

const name = "godox"

func New(settings *config.GodoxSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []goanalysis.Issue

	analyzer := &analysis.Analyzer{
		Name: name,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
		Run: func(pass *analysis.Pass) (any, error) {
			issues := runGodox(pass, settings)

			if len(issues) == 0 {
				return nil, nil
			}

			mu.Lock()
			resIssues = append(resIssues, issues...)
			mu.Unlock()

			return nil, nil
		},
	}

	return goanalysis.NewLinter(
		name,
		"Tool for detection of FIXME, TODO and other comment keywords",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithIssuesReporter(func(*linter.Context) []goanalysis.Issue {
		return resIssues
	}).WithLoadMode(goanalysis.LoadModeSyntax)
}

func runGodox(pass *analysis.Pass, settings *config.GodoxSettings) []goanalysis.Issue {
	var messages []godox.Message
	for _, file := range pass.Files {
		messages = append(messages, godox.Run(file, pass.Fset, settings.Keywords...)...)
	}

	if len(messages) == 0 {
		return nil
	}

	issues := make([]goanalysis.Issue, len(messages))

	for k, i := range messages {
		issues[k] = goanalysis.NewIssue(&result.Issue{
			Pos: token.Position{
				Filename: i.Pos.Filename,
				Line:     i.Pos.Line,
			},
			Text:       strings.TrimRight(i.Message, "\n"),
			FromLinter: name,
		}, pass)
	}

	return issues
}
