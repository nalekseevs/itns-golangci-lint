package processors

import (
	"github.com/nalekseevs/itns-golangci-lint/pkg/result"
)

const typeCheckName = "typecheck"

type Processor interface {
	Process(issues []result.Issue) ([]result.Issue, error)
	Name() string
	Finish()
}
