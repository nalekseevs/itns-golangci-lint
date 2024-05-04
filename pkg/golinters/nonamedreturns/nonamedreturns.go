package nonamedreturns

import (
	"github.com/firefart/nonamedreturns/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
)

func New(settings *config.NoNamedReturnsSettings) *goanalysis.Linter {
	a := analyzer.Analyzer

	var cfg map[string]map[string]any
	if settings != nil {
		cfg = map[string]map[string]any{
			a.Name: {
				analyzer.FlagReportErrorInDefer: settings.ReportErrorInDefer,
			},
		}
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		cfg,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
