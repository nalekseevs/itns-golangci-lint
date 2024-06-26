package gomoddirectives

import (
	"sync"

	"github.com/ldez/gomoddirectives"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
	"github.com/nalekseevs/itns-golangci-lint/pkg/lint/linter"
	"github.com/nalekseevs/itns-golangci-lint/pkg/result"
)

const linterName = "gomoddirectives"

func New(settings *config.GoModDirectivesSettings) *goanalysis.Linter {
	var issues []goanalysis.Issue
	var once sync.Once

	var opts gomoddirectives.Options
	if settings != nil {
		opts.ReplaceAllowLocal = settings.ReplaceLocal
		opts.ReplaceAllowList = settings.ReplaceAllowList
		opts.RetractAllowNoExplanation = settings.RetractAllowNoExplanation
		opts.ExcludeForbidden = settings.ExcludeForbidden
	}

	analyzer := &analysis.Analyzer{
		Name: goanalysis.TheOnlyAnalyzerName,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
		Run:  goanalysis.DummyRun,
	}

	return goanalysis.NewLinter(
		linterName,
		"Manage the use of 'replace', 'retract', and 'excludes' directives in go.mod.",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithContextSetter(func(lintCtx *linter.Context) {
		analyzer.Run = func(pass *analysis.Pass) (any, error) {
			once.Do(func() {
				results, err := gomoddirectives.Analyze(opts)
				if err != nil {
					lintCtx.Log.Warnf("running %s failed: %s: "+
						"if you are not using go modules it is suggested to disable this linter", linterName, err)
					return
				}

				for _, p := range results {
					issues = append(issues, goanalysis.NewIssue(&result.Issue{
						FromLinter: linterName,
						Pos:        p.Start,
						Text:       p.Reason,
					}, pass))
				}
			})

			return nil, nil
		}
	}).WithIssuesReporter(func(*linter.Context) []goanalysis.Issue {
		return issues
	}).WithLoadMode(goanalysis.LoadModeSyntax)
}
