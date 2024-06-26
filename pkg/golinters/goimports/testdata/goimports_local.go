//golangcitest:args -Egoimports
//golangcitest:config_path testdata/goimports_local.yml
package testdata

import (
	"fmt"

	"github.com/nalekseevs/itns-golangci-lint/pkg/config" // want "File is not `goimports`-ed with -local github.com/nalekseevs/itns-golangci-lint"
	"golang.org/x/tools/go/analysis"
)

func GoimportsLocalPrefixTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = analysis.Analyzer{}
}
