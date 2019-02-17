package durationcheck_test

import (
	"testing"

	"github.com/charithe/durationcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, durationcheck.Analyzer, "a")
}
