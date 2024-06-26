package nakedret

import (
	"github.com/alexkohler/nakedret/v2"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
)

func New(settings *config.NakedretSettings) *goanalysis.Linter {
	var maxLines int
	if settings != nil {
		maxLines = settings.MaxFuncLines
	}

	a := nakedret.NakedReturnAnalyzer(uint(maxLines))

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
