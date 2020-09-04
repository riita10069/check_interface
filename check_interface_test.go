package check_interface_test

import (
	"testing"

	"github.com/riita10069/check_interface"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(
		t,
		testdata,
		check_interface.Analyzer,
		"a",
		"b",
	//""
	)
}
