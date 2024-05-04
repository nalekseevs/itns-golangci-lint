package internal

import "github.com/nalekseevs/itns-golangci-lint/pkg/logutils"

// LinterLogger must be use only when the context logger is not available.
var LinterLogger = logutils.NewStderrLog(logutils.DebugKeyLinter)
