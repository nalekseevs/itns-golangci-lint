package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/nalekseevs/itns-golangci-lint/pkg/commands"
	"github.com/nalekseevs/itns-golangci-lint/pkg/exitcodes"
)

var (
	goVersion = "unknown"

	// Populated by goreleaser during build
	version = "unknown"
	commit  = "?"
	date    = ""
)

func main() {
	info := createBuildInfo()

	if err := commands.Execute(info); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed executing command with error: %v\n", err)
		os.Exit(exitcodes.Failure)
	}
}

func createBuildInfo() commands.BuildInfo {
	info := commands.BuildInfo{
		Commit:    commit,
		Version:   version,
		GoVersion: goVersion,
		Date:      date,
	}

	buildInfo, available := debug.ReadBuildInfo()
	if !available {
		return info
	}

	info.GoVersion = buildInfo.GoVersion

	if date != "" {
		return info
	}

	info.Version = buildInfo.Main.Version

	var revision string
	var modified string
	for _, setting := range buildInfo.Settings {
		// The `vcs.xxx` information is only available with `go build`.
		// This information is not available with `go install` or `go run`.
		switch setting.Key {
		case "vcs.time":
			info.Date = setting.Value
		case "vcs.revision":
			revision = setting.Value
		case "vcs.modified":
			modified = setting.Value
		}
	}

	if revision == "" {
		revision = "unknown"
	}

	if modified == "" {
		modified = "?"
	}

	if info.Date == "" {
		info.Date = "(unknown)"
	}

	info.Commit = fmt.Sprintf("(%s, modified: %s, mod sum: %q)", revision, modified, buildInfo.Main.Sum)

	return info
}
