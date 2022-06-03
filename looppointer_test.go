package looppointer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "simple")
	analysistest.Run(t, testdata, Analyzer, "fixed")
	analysistest.Run(t, testdata, Analyzer, "issue7")
	analysistest.Run(t, testdata, Analyzer, "nolint")
	analysistest.Run(t, testdata, Analyzer, "nested")
	analysistest.Run(t, testdata, Analyzer, "embedded-func-pointer")
}
