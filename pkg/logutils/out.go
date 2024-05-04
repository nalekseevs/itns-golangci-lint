package logutils

import (
	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

var (
	StdOut = color.Output // https://github.com/nalekseevs/itns-golangci-lint/issues/14
	StdErr = colorable.NewColorableStderr()
)
