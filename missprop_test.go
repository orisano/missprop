package missprop_test

import (
	"testing"

	"github.com/orisano/missprop"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, missprop.Analyzer, "a")
}