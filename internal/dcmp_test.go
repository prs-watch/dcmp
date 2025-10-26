// Black-box tests for dcmp command-line options.
package internal

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Test cases for each command-line option.
// Each case corresponds to a directory under testdata/ containing:
//   - a.txt: Before file
//   - b.txt: After file
//   - expected.txt: Expected output
var cases = []string{
	"no-options",
	"brief",
	"expand-tabs",
	"ignore-all-space",
	"ignore-blank-lines",
	"ignore-case",
	"ignore-matching-lines",
	"ignore-space-change",
	"report-identical-file",
	"strip-trailing-cr",
}

/*
captureStdout captures standard output during function execution.
Parameters:
  - f: Function to execute while capturing stdout

Returns the captured output as a string.
*/
func captureStdout(f func()) string {
	old := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	f()
	writer.Close()
	res, _ := io.ReadAll(reader)
	os.Stdout = old

	return string(res)
}

/*
TestExecute performs black-box testing of Execute function with various command-line options.
Each test case:
 1. Sets up parameters based on the option being tested
 2. Executes the function while capturing stdout
 3. Compares output against expected.txt
*/
func TestExecute(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			// Set up file paths
			bfpath := fmt.Sprintf("testdata/%s/a.txt", tc)
			afpath := fmt.Sprintf("testdata/%s/b.txt", tc)
			exppath := fmt.Sprintf("testdata/%s/expected.txt", tc)
			briefFlag := false
			colorMode := "always"
			expandTabsFlag := false
			ignoreAllSpaceFlag := false
			ignoreBlankLinesFlag := false
			ignoreCaseFlag := false
			ignoreMatchingLines := []string{}
			ignoreSpaceChangeFlag := false
			reportIdenticalFileFlag := false
			stripTrailingCRFlag := false

			if tc == "brief" {
				briefFlag = true
			}
			if tc == "expand-tabs" {
				expandTabsFlag = true
			}
			if tc == "ignore-all-space" {
				ignoreAllSpaceFlag = true
			}
			if tc == "ignore-blank-lines" {
				ignoreBlankLinesFlag = true
			}
			if tc == "ignore-case" {
				ignoreCaseFlag = true
			}
			if tc == "ignore-matching-lines" {
				ignoreMatchingLines = []string{"^#"}
			}
			if tc == "ignore-space-change" {
				ignoreSpaceChangeFlag = true
			}
			if tc == "report-identical-file" {
				reportIdenticalFileFlag = true
			}
			if tc == "strip-trailing-cr" {
				stripTrailingCRFlag = true
			}

			// run `internal.Execute`
			var res string
			res = captureStdout(func() {
				_ = Execute(
					bfpath,
					afpath,
					briefFlag,
					reportIdenticalFileFlag,
					ignoreBlankLinesFlag,
					ignoreCaseFlag,
					ignoreSpaceChangeFlag,
					ignoreAllSpaceFlag,
					colorMode,
					stripTrailingCRFlag,
					ignoreMatchingLines,
					expandTabsFlag,
				)
			})

			// assertion
			exp, _ := os.ReadFile(exppath)
			if res != string(exp) {
				t.Errorf("output mismatch:\ngot:\n%q\nwant:\n%q", res, string(exp))
			}
		})
	}
}
