package o11yspanname

import (
	"fmt"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var (
	Analyzer = &analysis.Analyzer{
		Name: "o11yspanname",
		Doc:  "Checks o11y StartSpan function calls",
		Run:  CheckO11ySpanName,
	}
)

func New() *goanalysis.Linter {

	return goanalysis.NewLinter(
		Analyzer.Name,
		Analyzer.Doc,
		[]*analysis.Analyzer{Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}

//func New(settings *config.StaticCheckSettings) *goanalysis.Linter {
//	cfg := internal.StaticCheckConfig(settings)
//
//	// `scconfig.Analyzer` is a singleton, then it's not possible to have more than one instance for all staticcheck "sub-linters".
//	// When we will merge the 4 "sub-linters", the problem will disappear: https://github.com/nalekseevs/itns-golangci-lint/issues/357
//	// Currently only stylecheck analyzer has a configuration in staticcheck.
//	scconfig.Analyzer.Run = func(_ *analysis.Pass) (any, error) {
//		return cfg, nil
//	}
//
//	analyzers := internal.SetupStaticCheckAnalyzers(stylecheck.Analyzers, internal.GetGoVersion(settings), cfg.Checks)
//
//	return goanalysis.NewLinter(
//		"stylecheck",
//		"Stylecheck is a replacement for golint",
//		analyzers,
//		nil,
//	).WithLoadMode(goanalysis.LoadModeTypesInfo)
//}

func CheckO11ySpanName(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, declaration := range file.Decls {
			if function, ok := declaration.(*ast.FuncDecl); ok {
				expectedSpanName := fmt.Sprintf(`"%s"`, genExpectedSpanName(pass, function))

				for _, l := range function.Body.List {
					if stmt, ok := l.(*ast.AssignStmt); ok {
						if lit, ok := extractSpanName(stmt); ok {
							if lit.Value != expectedSpanName {
								pass.Reportf(lit.Pos(), `bad span name: %s (expected: %s)`, lit.Value, expectedSpanName)
								//report.Report(pass, file.Doc, fmt.Sprintf(`bad span name: %s (expected: "%s")`, lit.Value, expectedSpanName))
								//PassReport(pass, lit, expectedSpanName)
							}
						}
					}
				}

			}

		}
	}
	return nil, nil
}

func genExpectedSpanName(pass *analysis.Pass, function *ast.FuncDecl) string {
	if function.Recv == nil {
		// Regular function
		return fmt.Sprintf("%s.%s", pass.Pkg.Name(), function.Name.Name)
	} else {
		// Method
		for _, recv := range function.Recv.List {
			if expr, ok := recv.Type.(*ast.StarExpr); ok {
				if ident, ok := expr.X.(*ast.Ident); ok {
					return fmt.Sprintf("(*%s).%s", ident.Name, function.Name.Name)
				}
			}
		}
	}

	return function.Name.Name
}

func extractSpanName(stmt *ast.AssignStmt) (*ast.BasicLit, bool) {
	if len(stmt.Rhs) == 1 {
		if callExpr, ok := stmt.Rhs[0].(*ast.CallExpr); ok {
			if len(callExpr.Args) != 2 {
				return nil, false
			}
			selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return nil, false
			}
			if selExpr.Sel.Name != "StartSpan" {
				return nil, false
			}

			if lit, ok := callExpr.Args[1].(*ast.BasicLit); ok {
				return lit, true
			}
		}
	}

	return nil, false
}
