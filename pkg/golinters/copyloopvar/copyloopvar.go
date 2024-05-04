package copyloopvar

import (
	"github.com/karamaru-alpha/copyloopvar"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
)

func New(settings *config.CopyLoopVarSettings) *goanalysis.Linter {
	a := copyloopvar.NewAnalyzer()

	var cfg map[string]map[string]any
	if settings != nil {
		cfg = map[string]map[string]any{
			a.Name: {
				"check-alias": settings.CheckAlias,
			},
		}
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		cfg,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
