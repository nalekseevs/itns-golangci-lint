package main

import (
	"github.com/nalekseevs/itns-golangci-lint/test/testdata_etc/unused_exported/lib"
)

func main() {
	lib.PublicFunc()
}
