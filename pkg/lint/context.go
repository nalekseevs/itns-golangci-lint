package lint

import (
	"context"
	"fmt"

	"github.com/nalekseevs/itns-golangci-lint/internal/pkgcache"
	"github.com/nalekseevs/itns-golangci-lint/pkg/config"
	"github.com/nalekseevs/itns-golangci-lint/pkg/exitcodes"
	"github.com/nalekseevs/itns-golangci-lint/pkg/fsutils"
	"github.com/nalekseevs/itns-golangci-lint/pkg/goanalysis/load"
	"github.com/nalekseevs/itns-golangci-lint/pkg/lint/linter"
	"github.com/nalekseevs/itns-golangci-lint/pkg/logutils"
)

type ContextBuilder struct {
	cfg *config.Config

	pkgLoader *PackageLoader

	fileCache *fsutils.FileCache
	pkgCache  *pkgcache.Cache

	loadGuard *load.Guard
}

func NewContextBuilder(cfg *config.Config, pkgLoader *PackageLoader,
	fileCache *fsutils.FileCache, pkgCache *pkgcache.Cache, loadGuard *load.Guard,
) *ContextBuilder {
	return &ContextBuilder{
		cfg:       cfg,
		pkgLoader: pkgLoader,
		fileCache: fileCache,
		pkgCache:  pkgCache,
		loadGuard: loadGuard,
	}
}

func (cl *ContextBuilder) Build(ctx context.Context, log logutils.Log, linters []*linter.Config) (*linter.Context, error) {
	pkgs, deduplicatedPkgs, err := cl.pkgLoader.Load(ctx, linters)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}

	if len(deduplicatedPkgs) == 0 {
		return nil, fmt.Errorf("%w: running `go mod tidy` may solve the problem", exitcodes.ErrNoGoFiles)
	}

	ret := &linter.Context{
		Packages: deduplicatedPkgs,

		// At least `unused` linters works properly only on original (not deduplicated) packages,
		// see https://github.com/nalekseevs/itns-golangci-lint/pull/585.
		OriginalPackages: pkgs,

		Cfg:       cl.cfg,
		Log:       log,
		FileCache: cl.fileCache,
		PkgCache:  cl.pkgCache,
		LoadGuard: cl.loadGuard,
	}

	return ret, nil
}
