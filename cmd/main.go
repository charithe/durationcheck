package main

import (
	"github.com/charithe/durationcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(durationcheck.Analyzer)
}
