package linter

import (
	"go/ast"

	"golang.org/x/tools/go/packages"

	"github.com/nalekseevs/itns-golangci-lint/internal/pkgcache"
	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/fsutils"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis/load"
	"github.com/nalekseevs/itns-golangci-lint/pkg/logutils"
)

type Context struct {
	// Packages are deduplicated (test and normal packages) packages
	Packages []*packages.Package

	// OriginalPackages aren't deduplicated: they contain both normal and test
	// version for each of packages
	OriginalPackages []*packages.Package

	Cfg       *config.Config
	FileCache *fsutils.FileCache
	Log       logutils.Log

	PkgCache  *pkgcache.Cache
	LoadGuard *load.Guard
}

func (c *Context) Settings() *config.LintersSettings {
	return &c.Cfg.LintersSettings
}

func (c *Context) ClearTypesInPackages() {
	for _, p := range c.Packages {
		clearTypes(p)
	}
	for _, p := range c.OriginalPackages {
		clearTypes(p)
	}
}

func clearTypes(p *packages.Package) {
	p.Types = nil
	p.TypesInfo = nil
	p.Syntax = []*ast.File{}
}
