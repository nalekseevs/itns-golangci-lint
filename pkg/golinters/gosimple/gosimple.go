package gosimple

import (
	"honnef.co/go/tools/simple"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
	"github.com/nalekseevs/itns-golangci-lint/pkg/golinters/internal"
)

func New(settings *config.StaticCheckSettings) *goanalysis.Linter {
	cfg := internal.StaticCheckConfig(settings)

	analyzers := internal.SetupStaticCheckAnalyzers(simple.Analyzers, internal.GetGoVersion(settings), cfg.Checks)

	return goanalysis.NewLinter(
		"gosimple",
		"Linter for Go source code that specializes in simplifying code",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
