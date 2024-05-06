package canonicalheader

import (
	"github.com/lasiar/canonicalheader"
	"golang.org/x/tools/go/analysis"

	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	a := canonicalheader.Analyzer

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}