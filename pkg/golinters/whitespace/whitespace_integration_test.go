package whitespace

import (
	"testing"

	"github.com/nalekseevs/itns-golangci-lint/test/testshared/integration"
)

func TestFix(t *testing.T) {
	integration.RunFix(t)
}

func TestFixPathPrefix(t *testing.T) {
	integration.RunFixPathPrefix(t)
}
