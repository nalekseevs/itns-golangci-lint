//golangcitest:args -Elll
//golangcitest:config_path testdata/lll_import.yml
//golangcitest:expected_exitcode 0
package testdata

import (
	anotherVeryLongImportAliasNameForTest "github.com/nalekseevs/itns-golangci-lint/internal/golinters"
	veryLongImportAliasNameForTest "github.com/nalekseevs/itns-golangci-lint/internal/golinters"
)

func LllMultiImport() {
	_ = veryLongImportAliasNameForTest.NewLLL(nil)
	_ = anotherVeryLongImportAliasNameForTest.NewLLL(nil)
}
